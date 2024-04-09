# Doc

## Deploying net

1. Launch net
2. Deploy contracts (optionally)
3. Get addresses using sfc abi

## Unlocking account

personal.unlockAccount(personal.listAccounts[0], 'password', 10000000);
"0x6ca548f6df5b540e72262e935b6fe3e72cdd68c9"

## Opera help

NAME:
   opera - the go-opera command line interface

   Copyright 2019-2021 The go-opera Authors

USAGE:
   opera [options] command [command options] [arguments...]
   
VERSION:
   1.1.3-rc.5-7f58df14-1712665667
   
COMMANDS:
   account                            Manage accounts
   attach                             Start an interactive JavaScript environment (connect to node)
   check                              Check blockchain
   checkconfig                        Checks configuration file
   console                            Start an interactive JavaScript environment
   db                                 A set of commands related to leveldb database
   dumpconfig                         Show configuration values
   export                             Export blockchain
   import                             Import a blockchain file
   js                                 Execute the specified JavaScript files
   license                            Display license information
   network                            network command [command options] [arguments...]
   snapshot                           A set of commands based on the snapshot
   validator                          Manage validators
   version                            Print version numbers
   wallet                             Manage Ethereum presale wallets
   help, h                            Shows a list of commands or help for one command
   
OPERA OPTIONS:
  --genesis value                     'path to genesis file' - sets the network genesis configuration.
  --genesis.allowExperimental         Allow to use experimental genesis file.
  --identity value                    Custom node name
  --datadir value                     Data directory for the databases and keystore (default: "/Users/posidoni/Library/Lachesis")
  --datadir.minfreedisk value         Minimum free disk space in MB, once reached triggers auto shut down (default = 8192 MB, 0 = disabled)
  --keystore value                    Directory for the keystore (default = inside the datadir)
  --usb                               Enable monitoring and management of USB hardware wallets
  --pcscdpath value                   Path to the smartcard daemon (pcscd) socket file
  --exitwhensynced.age value          Exits after synchronisation reaches the required age (default: 0s)
  --exitwhensynced.epoch value        Exits after synchronisation reaches the required epoch (default: 0)
  --lightkdf                          Reduce key-derivation RAM & CPU usage at some expense of KDF strength
  --config value                      TOML configuration file
  --validator.id value                ID of a validator to create events from (default: 0)
  --validator.pubkey value            Public key of a validator to create events from
  --validator.password value          Password to unlock validator private key
  --syncmode value                    Blockchain sync mode ("full" or "snap") (default: "full")
  --gcmode value                      Blockchain garbage collection mode ("light", "full", "archive") (default: "archive")
  --db.preset value                   DBs layout preset ('pbl-1' or 'ldb-1' or 'legacy-ldb' or 'legacy-pbl')
  --db.migration.mode value           MultiDB migration mode ('reformat' or 'rebuild')
  
TRANSACTION POOL OPTIONS:
  --txpool.locals value               Comma separated accounts to treat as locals (no flush, priority inclusion)
  --txpool.nolocals                   Disables price exemptions for locally submitted transactions
  --txpool.journal value              Disk journal for local transaction to survive node restarts (default: "transactions.rlp")
  --txpool.rejournal value            Time interval to regenerate the local transaction journal (default: 1h0m0s)
  --txpool.pricelimit value           Minimum gas price limit to enforce for acceptance into the pool (default: 1)
  --txpool.pricebump value            Price bump percentage to replace an already existing transaction (default: 10)
  --txpool.accountslots value         Minimum number of executable transaction slots guaranteed per account (default: 16)
  --txpool.globalslots value          Maximum number of executable transaction slots for all accounts (default: 5120)
  --txpool.accountqueue value         Maximum number of non-executable transaction slots permitted per account (default: 64)
  --txpool.globalqueue value          Maximum number of non-executable transaction slots for all accounts (default: 1024)
  --txpool.lifetime value             Maximum amount of time non-executable transaction are queued (default: 3h0m0s)
  
PERFORMANCE TUNING OPTIONS:
  --cache value                       Megabytes of memory allocated to internal caching (default: 3600)
  
ACCOUNT OPTIONS:
  --unlock value                      Comma separated list of accounts to unlock
  --password value                    Password file to use for non-interactive password input
  --signer value                      External signer (url or path to ipc file)
  --allow-insecure-unlock             Allow insecure account unlocking when account-related RPCs are exposed by http
  
API OPTIONS:
  --http                              Enable the HTTP-RPC server
  --http.addr value                   HTTP-RPC server listening interface (default: "localhost")
  --http.port value                   HTTP-RPC server listening port (default: 18545)
  --http.corsdomain value             Comma separated list of domains from which to accept cross origin requests (browser enforced)
  --http.vhosts value                 Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts '*' wildcard. (default: "localhost")
  --graphql                           Enable GraphQL on the HTTP-RPC server. Note that GraphQL can only be started if an HTTP server is started as well.
  --graphql.corsdomain value          Comma separated list of domains from which to accept cross origin requests (browser enforced)
  --graphql.vhosts value              Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts '*' wildcard. (default: "localhost")
  --http.api value                    API's offered over the HTTP-RPC interface
  --http.rpcprefix value              HTTP path path prefix on which JSON-RPC is served. Use '/' to serve on all paths.
  --ws                                Enable the WS-RPC server
  --ws.addr value                     WS-RPC server listening interface (default: "localhost")
  --ws.port value                     WS-RPC server listening port (default: 18546)
  --ws.api value                      API's offered over the WS-RPC interface
  --ws.origins value                  Origins from which to accept websockets requests
  --ws.rpcprefix value                HTTP path prefix on which JSON-RPC is served. Use '/' to serve on all paths.
  --ipcdisable                        Disable the IPC-RPC server
  --ipcpath value                     Filename for IPC socket/pipe within the datadir (explicit paths escape it)
  --rpc.gascap value                  Sets a cap on gas that can be used in ftm_call/estimateGas (0=infinite) (default: 50000000)
  --rpc.evmtimeout value              Sets a timeout used for eth_call (0=infinite) (default: 5s)
  --rpc.txfeecap value                Sets a cap on transaction fee (in FTM) that can be sent via the RPC APIs (0 = no cap) (default: 100)
  --rpc.timeout value                 Time limit for RPC calls execution (default: 5s)
  
CONSOLE OPTIONS:
  --jspath loadScript                 JavaScript root path for loadScript (default: ".")
  --exec value                        Execute JavaScript statement
  --preload value                     Comma separated list of JavaScript files to preload into the console
  
NETWORKING OPTIONS:
  --bootnodes value                   Comma separated enode URLs for P2P discovery bootstrap
  --port value                        Network listening port (default: 5050)
  --maxpeers value                    Maximum number of network peers (network disabled if set to 0) (default: 50)
  --maxpendpeers value                Maximum number of pending connection attempts (defaults used if set to 0) (default: 0)
  --nat value                         NAT port mapping mechanism (any|none|upnp|pmp|extip:<IP>) (default: "any")
  --nodiscover                        Disables the peer discovery mechanism (manual peer addition)
  --v5disc                            Enables the experimental RLPx V5 (Topic Discovery) mechanism
  --netrestrict value                 Restricts network communication to the given IP networks (CIDR masks)
  --iprestrict value                  Restricts network communication to the given IP addresses
  --privatenodes value                Comma separated enode URLs which must not be advertised as peers to public network
  --nodekey value                     P2P node key file
  --nodekeyhex value                  P2P node key as hex (for testing)
  
GAS PRICE ORACLE OPTIONS:
  
METRICS AND STATS OPTIONS:
  --metrics                              Enable metrics collection and reporting
  --metrics.expensive                    Enable expensive metrics collection and reporting
  --metrics.addr value                   Enable stand-alone metrics HTTP server listening interface (default: "127.0.0.1")
  --metrics.port value                   Metrics HTTP server listening port (default: 6060)
  --metrics.influxdb                     Enable metrics export/push to an external InfluxDB database
  --metrics.influxdb.endpoint value      InfluxDB API endpoint to report metrics to (default: "http://localhost:8086")
  --metrics.influxdb.database value      InfluxDB database name to push reported metrics to (default: "geth")
  --metrics.influxdb.username value      Username to authorize access to the database (default: "test")
  --metrics.influxdb.password value      Password to authorize access to the database (default: "test")
  --metrics.influxdb.tags value          Comma-separated InfluxDB tags (key/values) attached to all measurements (default: "host=localhost")
  --metrics.influxdbv2                   Enable metrics export/push to an external InfluxDB v2 database
  --metrics.influxdb.token value         Token to authorize access to the database (v2 only) (default: "test")
  --metrics.influxdb.bucket value        InfluxDB bucket name to push reported metrics to (v2 only) (default: "geth")
  --metrics.influxdb.organization value  InfluxDB organization name (v2 only) (default: "geth")
  --tracing                              Enable traces collection and reporting
  
TESTING OPTIONS:
  --fakenet value                     'n/N' - sets coinbase as fake n-th key from genesis of N validators.
  
LOGGING AND DEBUGGING OPTIONS:
  --verbosity value                   Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail (default: 3)
  --vmodule value                     Per-module verbosity: comma-separated list of <pattern>=<level> (e.g. eth/*=5,p2p=4)
  --log.json                          Format logs with JSON
  --log.backtrace value               Request a stack trace at a specific logging statement (e.g. "block.go:271")
  --log.debug                         Prepends log messages with call-site location (file and line number)
  --pprof                             Enable the pprof HTTP server
  --pprof.addr value                  pprof HTTP server listening interface (default: "127.0.0.1")
  --pprof.port value                  pprof HTTP server listening port (default: 6060)
  --pprof.memprofilerate value        Turn on memory profiling with the given rate (default: 524288)
  --pprof.blockprofilerate value      Turn on block profiling with the given rate (default: 0)
  --pprof.cpuprofile value            Write CPU profile to the given file
  --trace value                       Write execution trace to the given file
  
ALIASED (deprecated) OPTIONS:
  --nousb                             Disables monitoring for and managing USB hardware wallets (deprecated)
  --rpc                               Enable the HTTP-RPC server (deprecated and will be removed June 2021, use --http)
  --rpcaddr value                     HTTP-RPC server listening interface (deprecated and will be removed June 2021, use --http.addr) (default: "localhost")
  --rpcport value                     HTTP-RPC server listening port (deprecated and will be removed June 2021, use --http.port) (default: 18545)
  --rpccorsdomain value               Comma separated list of domains from which to accept cross origin requests (browser enforced) (deprecated and will be removed June 2021, use --http.corsdomain)
  --rpcvhosts value                   Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts '*' wildcard. (deprecated and will be removed June 2021, use --http.vhosts) (default: "localhost")
  --rpcapi value                      API's offered over the HTTP-RPC interface (deprecated and will be removed June 2021, use --http.api)
  
MISC OPTIONS:
  --help, -h                          show help
  --validatorsfile value              Path to validators file
  --network.type value                Network type ("mainnet", "testnet" or "devnet") (default: "devnet")
  

