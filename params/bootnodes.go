// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	//11 - 15
	"enode://64f706387bd0fd22852406ea6e1f376f6fce2ea6fad47e7ab25e8511081818a0f19871b7a4844f0660f57699f057de8764e0618251dc4e0b559b6f8c3b43d895@47.92.209.100:9215",
	"enode://07e866b8e19ba682f677ccb4d3cbd00768fbc7a08e24202a1903d998b2f25af8da38c0ad9d06987388f60458ab6cbdbdb8239dc355be8a277e52de5e9834284e@47.92.135.242:9215",
	"enode://99b0117d29ac053f977fc65a2bf3a322bd2a58357b30450e6e28ffd48deb262fd367fdc11c020e0a3bc5ab7a0c1ddf8adf9cc224233c53feca6c4b04ff4cec56@47.92.198.8:9215",
	"enode://c74a96e58a189ccd0d0c8a1699c6cbe04f1e648034ce7d261da26564ddbb1cfd2769df24bbd879390c46835eaf477cca40111ea91799d0df01c247e136521bce@47.92.207.147:9215",
	"enode://a04d48b7dde13ccba0012f6b9392ae3187a80701c661c47b6dfdaf6e8c522ae45708121ab6468e781e32515f4585f83de9313338cc18e429ed0b917ab28d14ff@47.92.75.213:9215",
	//mec
	//"enode://f3a5560955238902390b170d396263b994cd4d9b516f4230b81fe9c2252ddf2dc2f64521f2302d14efc59eeee1c1bf911179a0c5709ad9ef0281ad060ee5e063@221.221.234.95:9212",
	//16
	"enode://898bfe6fcb8e74852552192853035f52398815020500b9362abf00158be7c3301d0e3804b6085a69a004fec1bc4c9850c7517c6aff5b3c2dea8d575d4820fd44@47.92.220.57:9215",
	//31
	"enode://663cca9e279d0d569f7b1cfe119a36d5f091c7fc7f9c5972e40cfe976e7bc88481da68dcfcfe1ebeb138f38712180f28e03be6755605fc4642bc4d0e56aabe22@47.92.35.230:9215",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://7f8edaa4bc698a19666019f3858e35cc5b420907d61eb585aafd51a6a6f62e47f6cc5767890a9a38b51a3409520db566f04e60372b8523ecd9e770aa764ee603@47.92.224.044:9215",
	"enode://f54641b27755a6284ce6c3bf44410f6e83cc433604ac9cb8c862e4b41c5900ad0c4cd68c9fa51779732bc37ff9b34153eaa68b2d95b1d641db325fde211f6482@47.92.214.211:9215",
	"enode://1a998266cc8357581e1b2405f76039804d000fdbe71e58bc41f3facf2b9086d50303caf18be59bc2ed4bd283021a34e93e469ef8bae82540dff3c27f50467f1d@47.92.220.171:9215",
	"enode://676bf12d833e2c00e75ef0217980a6f2bfcd13e41cfaf4e47baab46b18042cd655c4b13e313dffa8e36b32bf6c7175bd29f2db7fb1d08465aa81e1859bf6e991@47.92.194.176:9215",
	"enode://676bf12d833e2c00e75ef0217980a6f2bfcd13e41cfaf4e47baab46b18042cd655c4b13e313dffa8e36b32bf6c7175bd29f2db7fb1d08465aa81e1859bf6e991@47.92.211.169:9215",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://7e9b4ad3a1b747cb78a53bdee18a96089a3540a4d87cdee0a600cfe7211ec986adc669beb3344eec8cdf083c93755fa90d3fd6cf6c9d4b4765e4fd93cfd4732e@47.92.100.100:9215", // 39
	"enode://81ee2b913f5c10695ddafc25264d8d68167aa1fc2993417c7b5d02f9b08e931a37820f1fd7c3e89a7e2fa6956a376f3d8bf9b2959bcaece76ceef83b26673b40@47.92.98.104:9215",  // 40
	"enode://d73b087c411957f55abbd2dbb9b2cfcfee41dec68dfb1bfcbb4710968ffc84cdeb9eb0cb261a6a904d5b8c9aa3e5cb98360fbbd0967b075ab73e186c75aac6bd@47.92.172.168:9215", // 2
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:30303",
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",
}
