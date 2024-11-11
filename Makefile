.PHONY:gen-mall-proto
gen-mall-proto:
    @cwgo server --server_name gomall --type RPC  --idl proto/*.proto -module github.com/youmie723/GoMall
