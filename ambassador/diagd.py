#!python

import sys

import json
import logging

import VERSION

from flask import Flask, render_template, request # Response, jsonify

from AmbassadorConfig import AmbassadorConfig
from envoy import EnvoyStats
from utils import RichStatus, SystemInfo, PeriodicTrigger

__version__ = VERSION.Version

logging.basicConfig(
    # filename=logPath,
    level=logging.DEBUG, # if appDebug else logging.INFO,
    format="%%(asctime)s diagd %s %%(levelname)s: %%(message)s" % __version__,
    datefmt="%Y-%m-%d %H:%M:%S"
)

estats = EnvoyStats()

dir_index = 1
health_checks = True
errors = 0

while sys.argv[dir_index].startswith('-'):
    arg = sys.argv[dir_index]
    dir_index += 1

    if arg == '--no-health':
        health_checks = False
    else:
        logging.error("unknown argument %s" % arg)
        errors += 1

if errors:
    sys.exit(errors)

if health_checks:
    logging.debug("Starting periodic updates")
    stats_updater = PeriodicTrigger(estats.update, period=5)

aconf = AmbassadorConfig(sys.argv[dir_index])

app = Flask(__name__)

@app.route('/ambassador/v0/check_alive', methods=[ 'GET' ])
def check_alive():
    age = estats.age()

    if age > 20:
        return "ambassador seems to have died (%d)" % age, 400
    else:
        return "ambassador liveness check OK (%d)" % age, 200

@app.route('/ambassador/v0/check_ready', methods=[ 'GET' ])
def check_ready():
    age = estats.age()

    if (age > 20) or (estats.stats['last_update'] == 0):
        return "ambassador not ready (%d)" % age, 400
    else:
        return "ambassador readiness check OK (%d)" % age, 200

@app.route('/ambassador/v0/diag/<path:source>', methods=[ 'GET' ])
def show_intermediate(source=None):
    logging.debug("getting intermediate for '%s'" % source)
    result = aconf.get_intermediate_for(source)

    method = request.args.get('method', None)
    resource = request.args.get('resource', None)

    # output = []

    # if result['sources']:
    #     for source in sorted(result['sources'], key=lambda x: "%s.%d" % (x['filename'], x['index'])):
    #         output.append("# %s[%d]" % (source['filename'], source['index']))
    #         output.append("---")
    #         output.append(source['yaml'])

    #     # link types back to Ambassador and Envoy docs
    #     # present cluster stats, too
    #     for type in [ 'listeners', 'filters', 'routes', 'clusters' ]:
    #         if result[type]:
    #             output.append("--------")
    #             output.append("%s:" % type)

    #             first = True

    #             for element in sorted(result[type], key=lambda x: x['_source']):
    #                 source = element['_source']
    #                 refs = element.get('_referenced_by', [])

    #                 output.append("# created by %s" % source)

    #                 if refs:
    #                     output.append("# referenced by %s" % ", ".join(sorted(refs)))

    #                 if type == 'clusters':
    #                     name = element['name']

    #                     cluster_stats = estats.cluster_stats(name)

    #                     if cluster_stats['valid']:
    #                         output.append("# cluster is %d%% healthy" % cluster_stats['healthy_percent'])
    #                     else:
    #                         output.append("# cluster health unknown: %s" % cluster_stats['reason'])

    #                 x = dict(**element)
    #                 del(x['_source'])

    #                 if '_referenced_by' in x:
    #                     del(x['_referenced_by'])

    #                 output.append("%s\n" % json.dumps(x, indent=4))

    #                 if not first:
    #                     output.append("")

    #                 first = False

    # return "\n".join(output)
    logging.debug(json.dumps(result, indent=4))

    cluster_names = [ x['name'] for x in result['clusters'] ]

    stats = { name: estats.cluster_stats(name) for name in cluster_names }
    print(stats)
    result['sources'].sort(key=lambda x: "%s.%d" % (x['filename'], x['index']))

    return render_template('diag.html', method=method, resource=resource, stats=stats, **result)

@app.template_filter('pretty_json')
def pretty_json(obj):
    if isinstance(obj, dict):
        obj = dict(**obj)

        if '_source' in obj:
            del(obj['_source'])

        if '_referenced_by' in obj:
            del(obj['_referenced_by'])

    return json.dumps(obj, indent=4, sort_keys=True)

app.run(host='127.0.0.1', port=8888, debug=True)
