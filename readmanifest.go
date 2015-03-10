package main;
// lifted from http://rosettacode.org/wiki/Walk_a_directory/Recursively#Go
import (
	"fmt"
	"os"
	"path/filepath"
	"encoding/json"
	"io/ioutil"
	"reflect"
)

/*
{
	"website": "https://www.odoo.com/page/billing",
	"description": "\nAccounting and Financial Management.\n====================================\n\nFinancial and accounting module that covers:\n--------------------------------------------\n    * General Accounting\n    * Cost/Analytic accounting\n    * Third party accounting\n    * Taxes management\n    * Budgets\n    * Customer and Supplier Invoices\n    * Bank statements\n    * Reconciliation process by partner\n\nCreates a dashboard for accountants that includes:\n--------------------------------------------------\n    * List of Customer Invoices to Approve\n    * Company Analysis\n    * Graph of Treasury\n\nProcesses like maintaining general ledgers are done through the defined Financial Journals (entry move line or grouping is maintained through a journal) \nfor a particular financial year and for preparation of vouchers there is a module named account_voucher.\n    ", "sequence": 100, "demo": ["demo/account_demo.xml", "project/project_demo.xml", "project/analytic_account_demo.xml", "demo/account_minimal.xml", "demo/account_invoice_demo.xml", "demo/account_bank_statement.xml", "account_unit_test.xml"], "depends": ["base_setup", "product", "analytic", "board", "edi", "report"], "auto_install": false, "data": ["security/account_security.xml", "security/ir.model.access.csv", "account_menuitem.xml", "report/account_invoice_report_view.xml", "report/account_entries_report_view.xml", "report/account_treasury_report_view.xml", "report/account_report_view.xml", "report/account_analytic_entries_report_view.xml", "wizard/account_move_bank_reconcile_view.xml", "wizard/account_use_model_view.xml", "account_installer.xml", "wizard/account_period_close_view.xml", "wizard/account_reconcile_view.xml", "wizard/account_unreconcile_view.xml", "wizard/account_statement_from_invoice_view.xml", "account_view.xml", "account_report.xml", "account_financial_report_data.xml", "wizard/account_report_common_view.xml", "wizard/account_invoice_refund_view.xml", "wizard/account_fiscalyear_close_state.xml", "wizard/account_chart_view.xml", "wizard/account_tax_chart_view.xml", "wizard/account_move_line_reconcile_select_view.xml", "wizard/account_open_closed_fiscalyear_view.xml", "wizard/account_move_line_unreconcile_select_view.xml", "wizard/account_vat_view.xml", "wizard/account_report_print_journal_view.xml", "wizard/account_report_general_journal_view.xml", "wizard/account_report_central_journal_view.xml", "wizard/account_subscription_generate_view.xml", "wizard/account_fiscalyear_close_view.xml", "wizard/account_state_open_view.xml", "wizard/account_journal_select_view.xml", "wizard/account_change_currency_view.xml", "wizard/account_validate_move_view.xml", "wizard/account_report_general_ledger_view.xml", "wizard/account_invoice_state_view.xml", "wizard/account_report_partner_balance_view.xml", "wizard/account_report_account_balance_view.xml", "wizard/account_report_aged_partner_balance_view.xml", "wizard/account_report_partner_ledger_view.xml", "wizard/account_reconcile_partner_process_view.xml", "wizard/account_automatic_reconcile_view.xml", "wizard/account_financial_report_view.xml", "wizard/pos_box.xml", "project/wizard/project_account_analytic_line_view.xml", "account_end_fy.xml", "account_invoice_view.xml", "data/account_data.xml", "data/data_account_type.xml", "data/configurable_account_chart.xml", "account_invoice_workflow.xml", "project/project_view.xml", "project/project_report.xml", "project/wizard/account_analytic_balance_report_view.xml", "project/wizard/account_analytic_cost_ledger_view.xml", "project/wizard/account_analytic_inverted_balance_report.xml", "project/wizard/account_analytic_journal_report_view.xml", "project/wizard/account_analytic_cost_ledger_for_journal_report_view.xml", "project/wizard/account_analytic_chart_view.xml", "partner_view.xml", "product_view.xml", "account_assert_test.xml", "ir_sequence_view.xml", "company_view.xml", "edi/invoice_action_data.xml", "account_bank_view.xml", "res_config_view.xml", "account_pre_install.yml", "views/report_vat.xml", "views/report_invoice.xml", "views/report_trialbalance.xml", "views/report_centraljournal.xml", "views/report_overdue.xml", "views/report_generaljournal.xml", "views/report_journal.xml", "views/report_salepurchasejournal.xml", "views/report_partnerbalance.xml", "views/report_agedpartnerbalance.xml", "views/report_partnerledger.xml", "views/report_partnerledgerother.xml", "views/report_financial.xml", "views/report_generalledger.xml", "project/views/report_analyticbalance.xml", "project/views/report_analyticjournal.xml", "project/views/report_analyticcostledgerquantity.xml", "project/views/report_analyticcostledger.xml", "project/views/report_invertedanalyticbalance.xml", "views/account.xml"], "demo_xml": [], "icon": "/base/static/description/icon.png", "category": "Accounting & Finance", "web": false, "init_xml": [], "name": "eInvoicing", "license": "AGPL-3", "author": "OpenERP SA", "update_xml": [], "summary": "", "application": false, "version": "8.0.1.1", "post_load": null, "test": ["test/account_test_users.yml", "test/account_customer_invoice.yml", "test/account_supplier_invoice.yml", "test/account_change_currency.yml", "test/chart_of_account.yml", "test/account_period_close.yml", "test/account_use_model.yml", "test/account_validate_account_move.yml", "test/test_edi_invoice.yml", "test/account_report.yml", "test/account_fiscalyear_close.yml"], "module_path": "/mnt/data/home/mdupont/experiments/odoo/addons/account", "qweb": ["static/src/xml/account_move_reconciliation.xml", "static/src/xml/account_move_line_quickadd.xml", "static/src/xml/account_bank_statement_reconciliation.xml"], "installable": true}
*/

type Load struct {}

type Manifest struct {
	website string
	data []string // array of file name
	demo []string // array of file names
	module_path string // file path
	icon string // relative file name
	init_xml []string
	test []string
	qweb []string
	author string
	summary string
	depends []string
	name string
	description string
	web bool
	post_load *Load
	version string
	demo_xml []string
	update_xml []string
	category string
	sequence int
	application bool
	auto_install bool
	installable bool
	license string
	external_dependancies	map[string][]string
}

func VisitPackageContents(fn string)  {
	file, e := ioutil.ReadFile(fn)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	// this is not loading
	// var m Manifest
	// err1 := json.Unmarshal(file, &m)
	// if err1 != nil {
	//	fmt.Printf("error1:%v\n", err1)
	//	return
	// }
	// fmt.Printf("Raw:%s:\n", m)

	var data map[string]interface {}
	err := json.Unmarshal(file, &data)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	for k, v := range data {
		fmt.Printf("Key:%s:\n", k)
		//fmt.Printf("pair:%s\t%s\n", k, v)

		switch t := v.(type) {
		case int:
			fmt.Printf("\tInteger: %v\n", t)
		case float64:
			fmt.Printf("\tFloat64: %v\n", t)
		case string:
			fmt.Printf("\tString: %v\n", t)
		case bool:
			fmt.Printf("\tString: %v\n", t)
		case nil:
			fmt.Printf("\tNil: %v\n", t)

		case map[string]interface{}:
			for i,n := range t {
				fmt.Printf("\t\tKV: %v\n", i)
				switch u := n.(type) {
				case []interface {}:
					for j,m := range u {
						fmt.Printf("\t\tItem: %v= %v\n", j, m)
					}
				}

			}

		case []interface {}:
			for i,n := range t {
				fmt.Printf("\t\tItem: %v= %v\n", i, n)
			}
		default:
			var r = reflect.TypeOf(t)
			fmt.Printf("\tOther:%v\n", r)
		}
	}





}

func VisitPackage(fp string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}
	if !!fi.IsDir() {
		return nil // not a file.  ignore.
	}
	matched, err := filepath.Match("__openerp__.json", fi.Name())
	if err != nil {
		fmt.Println(err) // malformed pattern
		return err       // this is fatal.
	}
	if matched {
		VisitPackageContents(fp)
	}
	return nil
}

func main() {
	filepath.Walk(".", VisitPackage)
}
