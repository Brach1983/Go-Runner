// *****************************************************************
// **************************CMS Runner****************************
// *****************************************************************

package main

import (
	"context"
	"time"
	"fmt"
	"database/sql"
	cdp "github.com/knq/chromedp"
    "github.com/knq/chromedp/runner"
    "github.com/knq/chromedp/cdp/network"
    _ "github.com/go-sql-driver/mysql"
	"os"
)


//***************************Craft CMS*****************************

//var username = ""
//var password = ""
//var title = " new post"
//var content = " Sample article"

//func Craftlogin() cdp.Tasks {
//  return cdp.Tasks{
//      cdp.Navigate(`http://`),



// **************************WordPress*****************************

var username = " admin"
var password = " Boost"
var title = " new post"
var content = " Sample article"

func wplogin() cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate(`http://brachwp.com/wp-login.php`),
		cdp.WaitVisible(`#user_login`, cdp.ByID),
		cdp.SendKeys(`#user_login`, username, cdp.ByID),
        cdp.SendKeys(`#user_pass`, password, cdp.ByID),
		cdp.Click("#wp-submit", cdp.ByID),
		cdp.Click("#wp-admin-bar-new-content", cdp.ByID),
		cdp.WaitVisible(`#wpbody-content`, cdp.ByID),
		cdp.SendKeys(`#title`, title, cdp.ByID),
		cdp.SendKeys(`#content`, content, cdp.ByID),
		cdp.Click("#publish", cdp.ByID),
		cdp.Sleep(4 * time.Second),
	}
}


// ***********************Joomla Primary***************************

var user = "brach"
var pass = "Boost"
var jtitle = " new post"
var jcontent = " sample article"

func jlogin(ctxt context.Context, c *cdp.CDP) {
    ct, cancel := context.WithDeadline(context.Background(), time.Now().Add(7*time.Second))
    defer cancel()
    c.Run(ctxt, cdp.Tasks{cdp.Navigate(`http://localhost:3000/administrator/index.php`)})
    c.Run(ct, cdp.Tasks{cdp.WaitVisible(`#element-box`, cdp.ByID)})
    c.Run(ctxt, cdp.Tasks{cdp.SendKeys(`#mod-login-username`, user, cdp.ByID)})
    c.Run(ctxt, cdp.Tasks{cdp.SendKeys(`#mod-login-password`, pass, cdp.ByID)})
    c.Run(ctxt, cdp.Tasks{cdp.Click("form#form-login>fieldset>div:nth-of-type(3)>div>div>button", cdp.ByQuery)})
    c.Run(ct, cdp.Tasks{cdp.WaitVisible(`#content`, cdp.ByID)})
    c.Run(ctxt, cdp.Tasks{cdp.Click("section#content>div>div>div>div>div>div>div>ul>li>a>span", cdp.ByQuery)})
    c.Run(ct, cdp.Tasks{cdp.WaitVisible(`#content`, cdp.ByID)})
    c.Run(ctxt, cdp.Tasks{cdp.SendKeys(`#jform_title`, jtitle, cdp.ByID)})
    c.Run(ctxt, cdp.Tasks{cdp.SendKeys(`#jform_articletext_ifr`, jcontent, cdp.ByID)})
    c.Run(ctxt, cdp.Tasks{cdp.Click("#toolbar-save", cdp.ByID)})
    c.Run(ct, cdp.Tasks{cdp.WaitVisible("body>nav>div>div>div>ul:nth-of-type(3)>li>a>span", cdp.ByQuery)})
    c.Run(ctxt, cdp.Tasks{cdp.Click("body>nav>div>div>div>ul:nth-of-type(3)>li>a>span", cdp.ByQuery)})
    c.Run(ct, cdp.Tasks{cdp.WaitVisible(`body>nav>div>div>div>ul:nth-of-type(3)>li>ul>li:nth-of-type(5)>a`, cdp.ByQuery)})
    c.Run(ctxt, cdp.Tasks{cdp.Click(`body>nav>div>div>div>ul:nth-of-type(3)>li>ul>li:nth-of-type(5)>a`, cdp.ByQuery)})

}


// **************************Joomla Secondary********************************

// var user = "brach"
// var pass = "Boost"
// var jtitle = " new post"
// var jcontent = " sample article"
//
// func jlogin() cdp.Tasks {
// 	return cdp.Tasks{
// 		cdp.Navigate(`http://localhost:3000/administrator/index.php`),
// 		cdp.WaitVisible(`#element-box`, cdp.ByID),
// 		cdp.SendKeys(`#mod-login-username`, user, cdp.ByID),
// 		cdp.SendKeys(`#mod-login-password`, pass, cdp.ByID),
// 		cdp.Click("form#form-login>fieldset>div:nth-of-type(3)>div>div>button", cdp.ByQuery),
// 		cdp.WaitVisible(`#content`, cdp.ByID),
// 		cdp.Click("section#content>div>div>div>div>div>div>div>ul>li>a>span", cdp.ByQuery),
// 		cdp.WaitVisible(`#content`, cdp.ByID),
// 		cdp.SendKeys(`#jform_title`, jtitle, cdp.ByID),
// 		cdp.SendKeys(`#jform_articletext_ifr`, jcontent, cdp.ByID),
// 		cdp.Click("#toolbar-save", cdp.ByID),
// 		cdp.WaitVisible("body>nav>div>div>div>ul:nth-of-type(3)>li>a>span", cdp.ByQuery),
// 		cdp.Click("body>nav>div>div>div>ul:nth-of-type(3)>li>a>span", cdp.ByQuery),
// 		cdp.WaitVisible(`body>nav>div>div>div>ul:nth-of-type(3)>li>ul>li:nth-of-type(5)>a`, cdp.ByQuery),
// 		cdp.Click("body>nav>div>div>div>ul:nth-of-type(3)>li>ul>li:nth-of-type(5)>a", cdp.ByQuery),
// 		cdp.Sleep(4 * time.Second),
// 	}
// }


func main() {
	var id string
	var accountId string
	var site string

	// create contextforberance
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := cdp.New(ctxt, cdp.WithRunnerOptions(runner.UserDataDir(".runnerConfig"), runner.Flag("--disable-web-security", true)))
	if err != nil {
	}

	db, err := sql.Open("mysql", "root:Boost@tcp(localhost:3306)/submitter_db")
	if err != nil {
		fmt.Println(err)
	}

	err2 := db.Ping()
	if err2 != nil {
		fmt.Println(err2)
	}

	rows, err := db.Query("SELECT id, account_id, site FROM cms_queue WHERE status = 0 LIMIT 1;")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&id, &accountId, &site)
		fmt.Println(id)
		fmt.Println(accountId)
		fmt.Println(site)

		// run task list
		if site == "wordpress" {
			err = c.Run(ctxt, wplogin())
			if err != nil {
				fmt.Println(err)
			}
		} else if site =="joomla" {
			jlogin(ctxt, c)
		}
		c.Run(ctxt, network.ClearBrowserCache())
		c.Run(ctxt, network.ClearBrowserCookies())
		fmt.Println("UPDATE cms_queue SET status=1 WHERE id = '" + id + "';")
		db.Exec("UPDATE cms_queue SET status=1 WHERE id = '" + id + "';")
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
	    }

	// wait for chrome to finish
	//err = c.Wait()
	//if err != nil {
	//    }

	os.RemoveAll(".runnerConfig")
	os.Exit(0)
}

