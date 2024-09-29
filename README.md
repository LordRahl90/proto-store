## Proto Store.

<p>This is meant to be a central repository for all my grpc related protofiles.<br /></p>
Uses `namely/protoc-all` to generate the necessary stubs

## Example

* `docker run -v $PWD/invoicer:/defs namely/protoc-all -f invoicer.proto -l go`

###TODO :
* Update the command to use makefile