# convert the plugin modules to go

# first convert the __openerp__.py to xml/json
import sys
#sys.path.append(
import openerp.modules.module as m

dkeys = ('external_dependencies', 
    'maintainer', 'sequence', 'demo', 'depends', 'images', 'category', 'web', 'init_xml', 'certificate', 'author', 'application', 'version', 'test', 'installable', 'website', 'description', 'auto_install', 'active', 'data', 'icon', 'name', 'license', 'update_xml', 'url', 'bootstrap', 'summary', 'complexity', 'demo_xml', 'post_load', 'css', 'qweb')

dkeys = {}
import json 
from os.path import join as opj
for mo in m.get_modules():
    m.load_openerp_module(mo)
    d = m.load_information_from_description_file(mo)
    mod_path = m.get_module_path(mo)
    f =  open(opj(mod_path, '__openerp__.json'),"w")
    #f.write(
    d['module_path']=mod_path
    json.dump(d,f)
    # print mo,
    # for data in d['data']:
    #     print data
    #     fullpath = m.get_module_resource(mo, data)
    #     #print res


