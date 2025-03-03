## 1. convertAddress

- request 
    
  ```bash
  grpcurl -plaintext -d '{
  "chain": "Kaspa",
  "publicKey": "03b3ac6cfddf9fcd6d2d1a5898b454ceade67cd773564c4058d7ed9a15c6a904dc"
  }' 127.0.0.1:8389 dapplink.utxo.WalletUtxoService.convertAddress
  ```
- response 
  
  ```bash
  {
  "code": "SUCCESS",
  "msg": "convert address success",
  "address": "kaspa:qypm8trvlh0elntd95d93x952n82menu6ae4vnzqtrt7mxs4c65sfhq9dh08yfg"
  } 
  
  ```


## 2. validate address
- request 
  ```bash
  grpcurl -plaintext -d '{
  "chain": "Kaspa",
  "address": "kaspa:qypm8trvlh0elntd95d93x952n82menu6ae4vnzqtrt7mxs4c65sfhq9dh08yfg"
  }' 127.0.0.1:8389 dapplink.utxo.WalletUtxoService.validAddress
  ```
  
- response 
  ```bash
  {
  "code": "SUCCESS",
  "msg": "kaspa address is valid",
  "valid": true
  }
  ```
  
 
## 3.get Account

- request 
  ```bash
  grpcurl -plaintext -d '{
  "chain": "Kaspa",
  "address": "kaspa:qqkqkzjvr7zwxxmjxjkmxxdwju9kjs6e9u82uh59z07vgaks6gg62v8707g73"
  }' 127.0.0.1:8389 dapplink.utxo.WalletUtxoService.getAccount
  ```
  
- response
  ```bash
  {
  "code":"SUCCESS",
  "msg":"get account success",
  "network":"kaspa-mainnet",
  "balance":"132655219913451.00"
  }
  ```
  
## 4. get utxo

- request 
  ```bash
  grpcurl -plaintext -d '{
  "chain": "Kaspa",
  "address": "kaspa:qqkqkzjvr7zwxxmjxjkmxxdwju9kjs6e9u82uh59z07vgaks6gg62v8707g73"
  }' 127.0.0.1:8389 dapplink.utxo.WalletUtxoService.getUnspentOutputs
  ```
  
- response
  ```bash
  {
  "code": "SUCCESS",
  "msg": "kaspa get unspent outputs success",
  "unspent_outputs": [
    {
      "tx_id": "3b5a9fb8e7e27f8f57c323e25d6811281997290c15821cd3e229be9ae3f907d3",
      "tx_hash_big_endian": "",
      "tx_output_n": "0",
      "script": "202c0b0a4c1f84e31b7234adb319ae970b6943592f0eae5e8513fcc476d0d211a5ac",
      "height": "54789012",
      "block_time": "",
      "address": "kaspa:qqkqkzjvr7zwxxmjxjkmxxdwju9kjs6e9u82uh59z07vgaks6gg62v8707g73",
      "unspent_amount": "5000000000",
      "value_hex": "",
      "confirmations": "0",
      "index": "0"
    },
    {
      "tx_id": "d90e2db4e8e16fc82a963df5f6cdf76097e219e850ca628a2d3b3122ce814485",
      "tx_hash_big_endian": "",
      "tx_output_n": "0",
      "script": "202c0b0a4c1f84e31b7234adb319ae970b6943592f0eae5e8513fcc476d0d211a5ac",
      "height": "18266887",
      "block_time": "",
      "address": "kaspa:qqkqkzjvr7zwxxmjxjkmxxdwju9kjs6e9u82uh59z07vgaks6gg62v8707g73",
      "unspent_amount": "123400000000",
      "value_hex": "",
      "confirmations": "0",
      "index": "0"
    }
  ]
  } 
  ```
  

## 5. get block by hash

- request 
  ```bash 
  grpcurl -plaintext -d '{
  "chain": "Kaspa",
  "hash": "a79d34cb4fcd656a7266846b3cf8ac45a50ea28a3b365299fc4b8f8680bf1d60"
  }' 127.0.0.1:8389 dapplink.utxo.WalletUtxoService.getBlockByHash
  ```
  
- response 
  ```bash
  {
  "code": "SUCCESS",
  "msg": "kaspa get block by block success",
  "height": "0",
  "hash": "a79d34cb4fcd656a7266846b3cf8ac45a50ea28a3b365299fc4b8f8680bf1d60",
  "tx_list": [
    {
      "hash": "64cf9f631d3529e04d1d0055d50b915bdd6fe6a1e78c5f1d18e52d193456ed34",
      "fee": "0",
      "vin": [],
      "vout": [
        {
          "address": "kaspa:qrk9decfnl4rayeegp6gd3tc6605zavclkpud5jp78axat5namppwt050d57j",
          "amount": "6543059132",
          "index": 0
        },
        {
          "address": "kaspa:qqc3a2j95vhn9jlq9d87mexyg7dwc0lvnyvzwypgwk9hx00h44krvlhf85g4q",
          "amount": "6540639132",
          "index": 1
        }
      ]
    },
    {
      "hash": "deb5831f41f0cf907660cecb45ef22f57b3446df14039216bee4d8c1d47bbf92",
      "fee": "0",
      "vin": [
        {
          "hash": "deb5831f41f0cf907660cecb45ef22f57b3446df14039216bee4d8c1d47bbf92",
          "index": 0,
          "amount": "0",
          "address": ""
        }
      ],
      "vout": [
        {
          "address": "kaspa:qpjz3p7ycfpu5yxms5zwk5jx4p0hfcc28hexdkycwp8f2d88jhjy7u40wj2z4",
          "amount": "500000000",
          "index": 0
        }
      ]
    },
    {
      "hash": "f7888af9365d6fcc2635d12c07201132c94540ae258a3ff831af57935bf9968c",
      "fee": "0",
      "vin": [
        {
          "hash": "f7888af9365d6fcc2635d12c07201132c94540ae258a3ff831af57935bf9968c",
          "index": 1,
          "amount": "0",
          "address": ""
        }
      ],
      "vout": [
        {
          "address": "kaspa:qypgfnwldtutklkvymm7je55e9wuqxt8kt3mf3fndj4sez7sl2use6g90e02wvh",
          "amount": "200592700420",
          "index": 0
        },
        {
          "address": "kaspa:qpqpyavkqnp60q6t4sfctz4yp3n0ct963z65rxkd5ft32vkehnd3wx8jqctr2",
          "amount": "22142440729015",
          "index": 1
        }
      ]
    },
    {
      "hash": "b8cb9693d57f23671d2d0415630e0f2422dc657dfbd017b1f49857f665c758e9",
      "fee": "0",
      "vin": [
        {
          "hash": "b8cb9693d57f23671d2d0415630e0f2422dc657dfbd017b1f49857f665c758e9",
          "index": 0,
          "amount": "0",
          "address": ""
        },
        {
          "hash": "b8cb9693d57f23671d2d0415630e0f2422dc657dfbd017b1f49857f665c758e9",
          "index": 1,
          "amount": "0",
          "address": ""
        }
      ],
      "vout": [
        {
          "address": "kaspa:qrf3rcpvkgkeqhsa7n4gndnvppj5edp7zh2mqj35mfc4rdm4ymkm7cdjrdl7y",
          "amount": "29950000",
          "index": 0
        },
        {
          "address": "kaspa:qpg2hksud4vr78znrjr3rlvaywj2rycvmmcespdnllhqlf6tk68qxlzswns2x",
          "amount": "623950000",
          "index": 1
        }
      ]
    }
  ]
  }
  ```
  
## 6. get block by number

- request 
  ```bash 
  grpcurl -plaintext -d '{
  "chain": "Kaspa",
  "height": "102820371"
  }' 127.0.0.1:8389 dapplink.utxo.WalletUtxoService.getBlockByNumber
  ```
  
- response 
  ```bash 
  {
  "code": "SUCCESS",
  "msg": "kaspa get block by block success",
  "height": "0",
  "hash": "a79d34cb4fcd656a7266846b3cf8ac45a50ea28a3b365299fc4b8f8680bf1d60",
  "tx_list": [
    {
      "hash": "deb5831f41f0cf907660cecb45ef22f57b3446df14039216bee4d8c1d47bbf92",
      "fee": "",
      "vin": [
        {
          "hash": "deb5831f41f0cf907660cecb45ef22f57b3446df14039216bee4d8c1d47bbf92",
          "index": 0,
          "amount": "0",
          "address": ""
        }
      ],
      "vout": [
        {
          "address": "kaspa:qpjz3p7ycfpu5yxms5zwk5jx4p0hfcc28hexdkycwp8f2d88jhjy7u40wj2z4",
          "amount": "500000000",
          "index": 0
        }
      ]
    },
    {
      "hash": "f7888af9365d6fcc2635d12c07201132c94540ae258a3ff831af57935bf9968c",
      "fee": "",
      "vin": [
        {
          "hash": "f7888af9365d6fcc2635d12c07201132c94540ae258a3ff831af57935bf9968c",
          "index": 1,
          "amount": "0",
          "address": ""
        }
      ],
      "vout": [
        {
          "address": "kaspa:qypgfnwldtutklkvymm7je55e9wuqxt8kt3mf3fndj4sez7sl2use6g90e02wvh",
          "amount": "200592700420",
          "index": 0
        },
        {
          "address": "kaspa:qpqpyavkqnp60q6t4sfctz4yp3n0ct963z65rxkd5ft32vkehnd3wx8jqctr2",
          "amount": "22142440729015",
          "index": 1
        }
      ]
    },
    {
      "hash": "b8cb9693d57f23671d2d0415630e0f2422dc657dfbd017b1f49857f665c758e9",
      "fee": "",
      "vin": [
        {
          "hash": "b8cb9693d57f23671d2d0415630e0f2422dc657dfbd017b1f49857f665c758e9",
          "index": 0,
          "amount": "0",
          "address": ""
        },
        {
          "hash": "b8cb9693d57f23671d2d0415630e0f2422dc657dfbd017b1f49857f665c758e9",
          "index": 1,
          "amount": "0",
          "address": ""
        }
      ],
      "vout": [
        {
          "address": "kaspa:qrf3rcpvkgkeqhsa7n4gndnvppj5edp7zh2mqj35mfc4rdm4ymkm7cdjrdl7y",
          "amount": "29950000",
          "index": 0
        },
        {
          "address": "kaspa:qpg2hksud4vr78znrjr3rlvaywj2rycvmmcespdnllhqlf6tk68qxlzswns2x",
          "amount": "623950000",
          "index": 1
        }
      ]
    }
  ]
  }
  ```
  
## 7. get tx by hash

- request
  ```bash
  grpcurl -plaintext -d '{
  "chain": "Kaspa",
  "hash": "56c8dd6f504ba2bf0777744c6c976ebc12c2b4845eaf2e71f557f3d954eb3464"
  }' 127.0.0.1:8389 dapplink.utxo.WalletUtxoService.getTxByHash
  ```
  
- response 
  ```bash
  {
  "code": "SUCCESS",
  "msg": "kaspa get block by hash success",
  "tx": {
    "hash": "56c8dd6f504ba2bf0777744c6c976ebc12c2b4845eaf2e71f557f3d954eb3464",
    "index": 0,
    "froms": [
      {
        "address": "kaspa:qpj3duwc54nf38l3cesq72kkj00sdpzpcxw8cm5phgskp5lgfv4n570cehyx2"
      },
      {
        "address": "kaspa:qpj3duwc54nf38l3cesq72kkj00sdpzpcxw8cm5phgskp5lgfv4n570cehyx2"
      },
      {
        "address": "kaspa:qpj3duwc54nf38l3cesq72kkj00sdpzpcxw8cm5phgskp5lgfv4n570cehyx2"
      }
    ],
    "tos": [
      {
        "address": "kaspa:qp43ncsq0zsreuppfutmucz9d78m7a486x7czmfwfvnze503pdgfwt90ea7ze"
      },
      {
        "address": "kaspa:qpj3duwc54nf38l3cesq72kkj00sdpzpcxw8cm5phgskp5lgfv4n570cehyx2"
      }
    ],
    "fee": "150000",
    "status": "Success",
    "values": [
      {
        "value": "39281723540"
      },
      {
        "value": "9999850000"
      }
    ],
    "type": 0,
    "height": "103081122",
    "brc20_address": "",
    "datetime": "2025-03-03T14:30:07+08:00"
  }
  }
  ```

## 8. get txs by address

- request
  ```bash 
  grpcurl -plaintext -d '{
  "chain": "Kaspa",
  "address": "kaspa:qqkqkzjvr7zwxxmjxjkmxxdwju9kjs6e9u82uh59z07vgaks6gg62v8707g73",
  "page": 1,
  "pagesize": 2
  }' 127.0.0.1:8389 dapplink.utxo.WalletUtxoService.getTxByAddress
  ```

- response 
  ```bash 
  {
  "code": "SUCCESS",
  "msg": "kaspa get tx by address success",
  "tx": [
    {
      "hash": "4d4f3f1d587987eab42db717a1fa28645be090f6ed661a136d7a3f577eb08afe",
      "index": 0,
      "froms": [
        {
          "address": "kaspa:qypmlgcxsm045d6ac9zm76tj9kvufqkzkykr0my7d3w5tyfnkxprp0g9lhr329c"
        }
      ],
      "tos": [
        {
          "address": "kaspa:qqkqkzjvr7zwxxmjxjkmxxdwju9kjs6e9u82uh59z07vgaks6gg62v8707g73"
        },
        {
          "address": "kaspa:qypmlgcxsm045d6ac9zm76tj9kvufqkzkykr0my7d3w5tyfnkxprp0g9lhr329c"
        }
      ],
      "fee": "10000",
      "status": "Success",
      "values": [
        {
          "value": "100000000"
        },
        {
          "value": "117366224"
        }
      ],
      "type": 0,
      "height": "88554711",
      "brc20_address": "",
      "datetime": "2024-09-16T15:51:58+08:00"
    },
    {
      "hash": "87780352a33b5bf5c73f5fb5619d01b0facec68258aae662df748e74fc0b16c1",
      "index": 0,
      "froms": [
        {
          "address": "kaspa:qpdh7qemtfc9ytflqaf9xdyzydgkrp57gl2frw3d6spnsrpvz3jtwfyt4wep0"
        },
        {
          "address": "kaspa:qpdh7qemtfc9ytflqaf9xdyzydgkrp57gl2frw3d6spnsrpvz3jtwfyt4wep0"
        }
      ],
      "tos": [
        {
          "address": "kaspa:qqkqkzjvr7zwxxmjxjkmxxdwju9kjs6e9u82uh59z07vgaks6gg62v8707g73"
        },
        {
          "address": "kaspa:qpdh7qemtfc9ytflqaf9xdyzydgkrp57gl2frw3d6spnsrpvz3jtwfyt4wep0"
        }
      ],
      "fee": "3220",
      "status": "Success",
      "values": [
        {
          "value": "110076594"
        },
        {
          "value": "99996780"
        }
      ],
      "type": 0,
      "height": "85727940",
      "brc20_address": "",
      "datetime": "2024-08-14T23:11:20+08:00"
    }
  ]
  }
  ```


