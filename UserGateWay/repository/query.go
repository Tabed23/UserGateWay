package repository
//Raw's Query to INSERT data into AWS RED SHIFT Database
var ( 
	Statdeposit = `INSERT INTO "dev"."public"."deposit" (id, coin, chain_detail, member_id, amount , status, request_time,	request_deadline, deposit_label, updated_time) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	Statlogin = `INSERT INTO "dev"."public"."login" ("member_id", "login_time","loginip","platform") VALUES($1, $2, $3, $4)`
	Statlogout = `INSERT INTO "dev"."public"."logout" (member_id,logout_time,logoutip ,platform) VALUES($1, $2, $3, $4)`
	Statreg = `INSERT INTO "dev"."public"."registration" (member_id, mobile, mobile_prefix, email ,status, register_ip, account_status, createdat, register_platform) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	Statwidthdrawl = `INSERT INTO "dev"."public"."withdrawals" (id, member_id, coin, status, amount, fee , address,	txid, submitedat, successedat, updatedat, chain_type, hash_info) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
)	

	