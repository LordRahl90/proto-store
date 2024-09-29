generate:
	#docker run --rm -v $PWD:/$(project) namely/protoc-all -f $(file).proto -l go
 	docker run -v $PWD/$(project):/defs namely/protoc-all -f $(file).proto -l go




# example `make generate project=invoicer file=invoicer`