package celestia_da_light_client

/*
celestia-da-light-client is a Celestia tracking light client for proving data inclusion.
It is primarily based on the ibc-go 07-tendermint light client, but has no ibc-go dependencies. 
It is similar to the 07-celestia light client whereas it uses the data hash for the
consensus state root instead of the app hash and re-uses celestia-core's share/row proof validation.
Unlike 07-celestia, it does not use 07-tendermint's client state and does not require ibc-go
core modifications (as of this writing).
*/
