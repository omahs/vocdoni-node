package vochain

// Testnet Genesis File for Vocdoni KISS v1
const (
	ReleaseGenesis1 = `
   {
      "genesis_time":"2020-09-15T15:38:33.672114557Z",
      "chain_id":"vocdoni-release-06",
      "consensus_params":{
         "block":{
            "max_bytes":"22020096",
            "max_gas":"-1",
            "time_iota_ms":"10000"
         },
         "evidence":{
            "max_age_num_blocks":"100000",
            "max_age_duration":"10000"
         },
         "validator":{
            "pub_key_types":[
               "ed25519"
            ]
         }
      },
      "validators":[
         {
            "address":"6DB4FEE1D370907B31196B493714FC0F45C62DED",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"R7U+HxyTrvlXccEm1sc80ww83Fpp4xg247nmpjmkYTc="
            },
            "power":"10",
            "name":"miner1"
         },
         {
            "address":"71AA2FEFA96447BC5AEF9FD928F3F8ED57E695CF",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"ixI91P+MP1jiVIy1JwQqwRdZIZxsVI0WrytAzohMGCk="
            },
            "power":"10",
            "name":"miner2"
         },
         {
            "address":"AA9CC01B46BDD1AC9E2197BB9B84993CCDF880B2",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"H6oEMFrNFeQemr9Kgxjq/wVk1kZQ1VE/J1wVnVJ+K9I="
            },
            "power":"10",
            "name":"miner3"
         },
         {
            "address":"314D17BBE991FBD3D234E5C62CFD5D0717123C95",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"FLEg/pgdF4dZ060mved/z99p/EJePu9kSsyLnrsRNC0="
            },
            "power":"10",
            "name":"miner4"
         },
         {
            "address":"34B048A4A720E6B3918CF8B75CF12555080465E5",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"aF/+WaNs5tknRMRpTPO49TJZLmDctO+JH8uckE5fTNU="
            },
            "power":"10",
            "name":"miner5"
         }
      ],
      "app_hash":"",
      "app_state":{
         "validators":[
            {
               "address":"6DB4FEE1D370907B31196B493714FC0F45C62DED",
               "pub_key":{
                  "type":"tendermint/PubKeyEd25519",
                  "value":"R7U+HxyTrvlXccEm1sc80ww83Fpp4xg247nmpjmkYTc="
               },
               "power":"10",
               "name":"miner1"
            },
            {
               "address":"71AA2FEFA96447BC5AEF9FD928F3F8ED57E695CF",
               "pub_key":{
                  "type":"tendermint/PubKeyEd25519",
                  "value":"ixI91P+MP1jiVIy1JwQqwRdZIZxsVI0WrytAzohMGCk="
               },
               "power":"10",
               "name":"miner2"
            },
            {
               "address":"AA9CC01B46BDD1AC9E2197BB9B84993CCDF880B2",
               "pub_key":{
                  "type":"tendermint/PubKeyEd25519",
                  "value":"H6oEMFrNFeQemr9Kgxjq/wVk1kZQ1VE/J1wVnVJ+K9I="
               },
               "power":"10",
               "name":"miner3"
            },
            {
               "address":"314D17BBE991FBD3D234E5C62CFD5D0717123C95",
               "pub_key":{
                  "type":"tendermint/PubKeyEd25519",
                  "value":"FLEg/pgdF4dZ060mved/z99p/EJePu9kSsyLnrsRNC0="
               },
               "power":"10",
               "name":"miner4"
            },
            {
               "address":"34B048A4A720E6B3918CF8B75CF12555080465E5",
               "pub_key":{
                  "type":"tendermint/PubKeyEd25519",
                  "value":"aF/+WaNs5tknRMRpTPO49TJZLmDctO+JH8uckE5fTNU="
               },
               "power":"10",
               "name":"miner5"
            }
         ],
         "oracles":[
            "0xc2e396d6e6ae9b12551f0c6111f9766bec926bfe",
            "0x1a361c26e04a33effbf3bd8617b1e3e0aa6b704f"
         ]
      }
   }
 `
	DevelopmentGenesis1 = `
{
   "genesis_time":"2020-09-06T10:29:50.512370579Z",
   "chain_id":"vocdoni-development-17",
   "consensus_params":{
      "block":{
         "max_bytes":"22020096",
         "max_gas":"-1",
         "time_iota_ms":"10000"
      },
      "evidence":{
         "max_age_num_blocks":"100000",
         "max_age_duration":"10000"
      },
      "validator":{
         "pub_key_types":[
            "ed25519"
         ]
      }
   },
   "validators":[
      {
         "address":"5C69093136E0CB84E5CFA8E958DADB33C0D0CCCF",
         "pub_key":{
            "type":"tendermint/PubKeyEd25519",
            "value":"mXc5xXTKgDSYcy1lBCT1Ag7Lh1nPWHMa/p80XZPzAPY="
         },
         "power":"10",
         "name":"miner0"
      },
      {
         "address":"2E1B244B84E223747126EF621C022D5CEFC56F69",
         "pub_key":{
            "type":"tendermint/PubKeyEd25519",
            "value":"gaf2ZfdxpoielRXDXyBcMxkdzywcE10WsvLMe1K62UY="
         },
         "power":"10",
         "name":"miner1"
      },
      {
         "address":"4EF00A8C18BD472167E67F28694F31451A195581",
         "pub_key":{
            "type":"tendermint/PubKeyEd25519",
            "value":"dZXMBiQl4s0/YplfX9iMnCWonJp2gjrFHHXaIwqqtmc="
         },
         "power":"10",
         "name":"miner2"
      },
      {
         "address":"ECCC09A0DF8F4E5554A9C58F634E9D6AFD5F1598",
         "pub_key":{
            "type":"tendermint/PubKeyEd25519",
            "value":"BebelLYe4GZKwy9IuXCyBTySxQCNRrRoi1DSvAf6QxE="
         },
         "power":"10",
         "name":"miner3"
      },
      {
         "address":"3272B3046C31D87F92E26D249B97CC144D835DA6",
         "pub_key":{
            "type":"tendermint/PubKeyEd25519",
            "value":"hUz8jCePBfG4Bi9s13IdleWq5MZ5upe03M+BX2ah7c4="
         },
         "power":"10",
         "name":"miner4"
      },
      {
         "address":"05BA8FCBEA4A4EDCFD49081B42CA3F9ED13246C1",
         "pub_key":{
            "type":"tendermint/PubKeyEd25519",
            "value":"1FGNernnvg4QpV7psYFQPeIFZJm32yN1SjZULbliidg="
         },
         "power":"10",
         "name":"miner5"
      }
   ],
   "app_hash":"",
   "app_state":{
      "validators":[
         {
            "address":"5C69093136E0CB84E5CFA8E958DADB33C0D0CCCF",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"mXc5xXTKgDSYcy1lBCT1Ag7Lh1nPWHMa/p80XZPzAPY="
            },
            "power":"10",
            "name":"miner0"
         },
         {
            "address":"2E1B244B84E223747126EF621C022D5CEFC56F69",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"gaf2ZfdxpoielRXDXyBcMxkdzywcE10WsvLMe1K62UY="
            },
            "power":"10",
            "name":"miner1"
         },
         {
            "address":"4EF00A8C18BD472167E67F28694F31451A195581",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"dZXMBiQl4s0/YplfX9iMnCWonJp2gjrFHHXaIwqqtmc="
            },
            "power":"10",
            "name":"miner2"
         },
         {
            "address":"ECCC09A0DF8F4E5554A9C58F634E9D6AFD5F1598",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"BebelLYe4GZKwy9IuXCyBTySxQCNRrRoi1DSvAf6QxE="
            },
            "power":"10",
            "name":"miner3"
         },
         {
            "address":"3272B3046C31D87F92E26D249B97CC144D835DA6",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"hUz8jCePBfG4Bi9s13IdleWq5MZ5upe03M+BX2ah7c4="
            },
            "power":"10",
            "name":"miner4"
         },
         {
            "address":"05BA8FCBEA4A4EDCFD49081B42CA3F9ED13246C1",
            "pub_key":{
               "type":"tendermint/PubKeyEd25519",
               "value":"1FGNernnvg4QpV7psYFQPeIFZJm32yN1SjZULbliidg="
            },
            "power":"10",
            "name":"miner5"
         }
      ],
      "oracles":[
         "0xb926be24A9ca606B515a835E91298C7cF0f2846f",
         "0x2f4ed2773dcf7ad0ec15eb84ec896f4eebe0e08a"
      ]
   }
}
`
)
