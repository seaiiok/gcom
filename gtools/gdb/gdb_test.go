package gdb

import (
	"fmt"
	"testing"

	_ "github.com/alexbrainman/odbc"
)

func TestDB(t *testing.T) {
	dsn := fmt.Sprintf("driver={sql server};server=%s;port=%d;uid=%s;pwd=%s;database=%s;encrypt=disable", "192.168.1.18", 1433, "sa", "kmtSoft12345678", "ifixsvr")
	err := New("odbc", dsn)
	t.Log(err)

// 	createtable := `
// 	CREATE TABLE [dbo].[iFixsvr_JFOffline_info] (
// 		[id] [decimal](18, 0) IDENTITY (1, 1) NOT NULL ,
// 		[FilePath] [char] (10) COLLATE Chinese_PRC_CI_AS NULL ,
// 		[MD5] [char] (10) COLLATE Chinese_PRC_CI_AS NULL 
// 	) ON [PRIMARY]
// `
// 	err = Exec(createtable)
// 	t.Log(err)

// 		inserts := `
// 	insert into [iFixsvr_JFOffline_info] (FilePath,MD5) values (?,?)
// 	`
// 		re := []string{"1", "2"}

// 		res := make([][]string, 0)
// 		res = append(res, re)

// 		err = Inserts(inserts, res)
// 		t.Log(err)

// 		querys:=`select * from [iFixsvr_JFOffline_info]`

// 		qres,err:=Querys(querys)
// 		t.Log(err)
// 		t.Log(qres)

		// EXISTS (SELECT *
        //     FROM [iFixsvr_JFOffline_info]
		//     WHERE MD5 = '2')
	exist:=`SELECT * FROM [iFixsvr_JFOffline_info] WHERE MD5 = '2'`	
		res,err:=Exist(exist)
		t.Log(res,err)
}
