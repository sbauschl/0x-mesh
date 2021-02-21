package ethereum

import (
	"fmt"

	"github.com/0xProject/0x-mesh/constants"
	"github.com/ethereum/go-ethereum/common"
)

// ContractAddresses maps a contract's name to it's Ethereum address
type ContractAddresses struct {
	ERC20Proxy          common.Address `json:"erc20Proxy"`
	ERC721Proxy         common.Address `json:"erc721Proxy"`
	ERC1155Proxy        common.Address `json:"erc1155Proxy"`
	Exchange            common.Address `json:"exchange"`
	Coordinator         common.Address `json:"coordinator"`
	CoordinatorRegistry common.Address `json:"coordinatorRegistry"`
	DevUtils            common.Address `json:"devUtils"`
	WETH9               common.Address `json:"weth9"`
	ZRXToken            common.Address `json:"zrxToken"`
	ChaiBridge          common.Address `json:"chaiBridge"`
	ChaiToken           common.Address `json:"chaiToken"`
	MaximumGasPrice     common.Address `json:"maximumGasPrice"`
}

// GanacheAddresses The addresses that the 0x contracts were deployed to on the Ganache snapshot (chainID = 1337).
var GanacheAddresses = ganacheAddresses()

// NewContractAddressesForChainID The default contract addresses for the standard chainIDs.
func NewContractAddressesForChainID(chainID int) (ContractAddresses, error) {
	switch chainID {
	case 137:
		return ContractAddresses{
			ERC20Proxy:          common.HexToAddress("0x411b0bcf1b6ea88cb7229558c89994a2449c302c"),
			ERC721Proxy:         common.HexToAddress("0x58807bad0b376efc12f5ad86aac70e78ed67deae"),
			Exchange:            common.HexToAddress("0xfede379e48c873c75f3cc0c81f7c784ad730a8f7"),
			ERC1155Proxy:        common.HexToAddress("0x207fa8df3a17d96ca7ea4f2893fcdcb78a304101"),
			Coordinator:         common.HexToAddress("0xe45880fea5442a1914ddaa2e12411889f2b47a7d"),
			CoordinatorRegistry: common.HexToAddress("0x12824d0ba9972458426d9000ab61ee906c8bd87e"),
			DevUtils:            common.HexToAddress("0x68724fda24ce973ad8ad75a333ca1d7e046e1aa0"),
			WETH9:               common.HexToAddress("0x0849b784a770027817301cc46cef753e9016b604"),
			ZRXToken:            common.HexToAddress("0x03e1d56dfe6f4000ae38c7f8d002a6b0af55331c"),
			ChaiBridge:          common.HexToAddress("0x5638a4b19f121adc4436de3f0e845173b33b594c"),
			ChaiToken:           common.HexToAddress("0x5b5e11e4818cceba3e82ca9b97cd0ab80be75ad3"),
			MaximumGasPrice:     common.HexToAddress("0xe2bfd35306495d11e3c9db0d8de390cda24563cf"),
		}, nil
	case 1:
		return ContractAddresses{
			ERC20Proxy:          common.HexToAddress("0x0b47076aaa5246411458fcf85494f41bbfdb8470"),
			ERC721Proxy:         common.HexToAddress("0xff7ca10af37178bdd056628ef42fd7f799fac77c"),
			Exchange:            common.HexToAddress("0x533dc89624dcc012c7323b41f286bd2df478800b"),
			ERC1155Proxy:        common.HexToAddress("0x53d791f18155c211ff8b58671d0f7e9b50e596ad"),
			Coordinator:         common.HexToAddress("0x6669d66979f729445826fee33021090599ad7bd2"),
			CoordinatorRegistry: common.HexToAddress("0x6f5b9e0456c4849224c7b59dc15f05c48641c4e3"),
			DevUtils:            common.HexToAddress("0x7a2d89c4cb4b28b9cef9f269d48b3aecf0f549b7"),
			WETH9:               common.HexToAddress("0x5b5e11e4818cceba3e82ca9b97cd0ab80be75ad3"),
			ZRXToken:            common.HexToAddress("0xe01ac7c9eb19c63b063134ed2bb5eb7dcc847be9"),
			ChaiBridge:          common.HexToAddress("0x5638a4b19f121adc4436de3f0e845173b33b594c"),
			ChaiToken:           common.HexToAddress("0x5b5e11e4818cceba3e82ca9b97cd0ab80be75ad3"),
			MaximumGasPrice:     common.HexToAddress("0xe2bfd35306495d11e3c9db0d8de390cda24563cf"),
		}, nil
	case 3:
		return ContractAddresses{
			ERC20Proxy:          common.HexToAddress("0xb1408f4c245a23c31b98d2c626777d4c0d766caa"),
			ERC721Proxy:         common.HexToAddress("0xe654aac058bfbf9f83fcaee7793311dd82f6ddb4"),
			Exchange:            common.HexToAddress("0xfb2dd2a1366de37f7241c83d47da58fd503e2c64"),
			ERC1155Proxy:        common.HexToAddress("0x19bb6caa3bc34d39e5a23cedfa3e6c7e7f3c931d"),
			Coordinator:         common.HexToAddress("0x6ff734d96104965c9c1b0108f83abc46e6e501df"),
			CoordinatorRegistry: common.HexToAddress("0x403cc23e88c17c4652fb904784d1af640a6722d9"),
			DevUtils:            common.HexToAddress("0xb1a3d901bad1df7d710fc8d008db7cdd6bbbffe6"),
			WETH9:               common.HexToAddress("0xc778417e063141139fce010982780140aa0cd5ab"),
			ZRXToken:            common.HexToAddress("0xff67881f8d12f372d91baae9752eb3631ff0ed00"),
			ChaiBridge:          common.HexToAddress("0x0000000000000000000000000000000000000000"),
			ChaiToken:           common.HexToAddress("0x0000000000000000000000000000000000000000"),
			MaximumGasPrice:     common.HexToAddress("0x407b4128e9ecad8769b2332312a9f655cb9f5f3a"),
		}, nil
	case 4:
		return ContractAddresses{
			ERC20Proxy:          common.HexToAddress("0x2f5ae4f6106e89b4147651688a92256885c5f410"),
			ERC721Proxy:         common.HexToAddress("0x7656d773e11ff7383a14dcf09a9c50990481cd10"),
			ERC1155Proxy:        common.HexToAddress("0x19bb6caa3bc34d39e5a23cedfa3e6c7e7f3c931d"),
			Exchange:            common.HexToAddress("0x198805e9682fceec29413059b68550f92868c129"),
			Coordinator:         common.HexToAddress("0x70c5385ee5ee4629ef72abd169e888c8b4a12238"),
			CoordinatorRegistry: common.HexToAddress("0x1084b6a398e47907bae43fec3ff4b677db6e4fee"),
			DevUtils:            common.HexToAddress("0xb1a3d901bad1df7d710fc8d008db7cdd6bbbffe6"),
			WETH9:               common.HexToAddress("0xc778417e063141139fce010982780140aa0cd5ab"),
			ZRXToken:            common.HexToAddress("0x8080c7e4b81ecf23aa6f877cfbfd9b0c228c6ffa"),
			ChaiBridge:          common.HexToAddress("0x0000000000000000000000000000000000000000"),
			ChaiToken:           common.HexToAddress("0x0000000000000000000000000000000000000000"),
			MaximumGasPrice:     common.HexToAddress("0x47697b44bd89051e93b4d5857ba8e024800a74ac"),
		}, nil
	case 42:
		return ContractAddresses{
			ERC20Proxy:          common.HexToAddress("0xf1ec01d6236d3cd881a0bf0130ea25fe4234003e"),
			ERC721Proxy:         common.HexToAddress("0x2a9127c745688a165106c11cd4d647d2220af821"),
			Exchange:            common.HexToAddress("0x4eacd0af335451709e1e7b570b8ea68edec8bc97"),
			ERC1155Proxy:        common.HexToAddress("0x64517fa2b480ba3678a2a3c0cf08ef7fd4fad36f"),
			Coordinator:         common.HexToAddress("0xd29e59e51e8ab5f94121efaeebd935ca4214e257"),
			CoordinatorRegistry: common.HexToAddress("0x09fb99968c016a3ff537bf58fb3d9fe55a7975d5"),
			DevUtils:            common.HexToAddress("0xb1a3d901bad1df7d710fc8d008db7cdd6bbbffe6"),
			WETH9:               common.HexToAddress("0xd0a1e359811322d97991e03f863a0c30c2cf029c"),
			ZRXToken:            common.HexToAddress("0x2002d3812f58e35f0ea1ffbf80a75a38c32175fa"),
			ChaiBridge:          common.HexToAddress("0x0000000000000000000000000000000000000000"),
			ChaiToken:           common.HexToAddress("0x0000000000000000000000000000000000000000"),
			MaximumGasPrice:     common.HexToAddress("0x67a094cf028221ffdd93fc658f963151d05e2a74"),
		}, nil
	case 1337:
		return ganacheAddresses(), nil
	default:
		return ContractAddresses{}, fmt.Errorf("Cannot create contract addresses for non-standard chainID")
	}
}

func ValidateContractAddressesForChainID(chainID int, addresses ContractAddresses) error {
	if chainID == 1 {
		return fmt.Errorf("cannot add contract addresses for chainID 1: addresses for mainnet are hard-coded and cannot be changed")
	}
	if addresses.Exchange == constants.NullAddress {
		return fmt.Errorf("cannot add contract addresses for chain ID %d: Exchange address is required", chainID)
	}
	if addresses.DevUtils == constants.NullAddress {
		return fmt.Errorf("cannot add contract addresses for chain ID %d: DevUtils address is required", chainID)
	}
	if addresses.ERC20Proxy == constants.NullAddress {
		return fmt.Errorf("cannot add contract addresses for chain ID %d: ERC20Proxy address is required", chainID)
	}
	if addresses.ERC721Proxy == constants.NullAddress {
		return fmt.Errorf("cannot add contract addresses for chain ID %d: ERC721Proxy address is required", chainID)
	}
	if addresses.ERC1155Proxy == constants.NullAddress {
		return fmt.Errorf("cannot add contract addresses for chain ID %d: ERC1155Proxy address is required", chainID)
	}
	// TODO(albrow): Uncomment this if we re-add coordinator support.
	// if addresses.CoordinatorRegistry == constants.NullAddress {
	// 	return fmt.Errorf("cannot add contract addresses for chain ID %d: CoordinatorRegistry address is required", chainID)
	// }
	return nil
}

// ganacheAddresses Returns the addresses of the deployed contracts on the Ganache snapshot. This
// function allows these addresses to only be defined in one place.
func ganacheAddresses() ContractAddresses {
	return ContractAddresses{
		ERC20Proxy:          common.HexToAddress("0x1dc4c1cefef38a777b15aa20260a54e584b16c48"),
		ERC721Proxy:         common.HexToAddress("0x1d7022f5b17d2f8b695918fb48fa1089c9f85401"),
		ERC1155Proxy:        common.HexToAddress("0x6a4a62e5a7ed13c361b176a5f62c2ee620ac0df8"),
		Exchange:            common.HexToAddress("0x48bacb9266a570d521063ef5dd96e61686dbe788"),
		Coordinator:         common.HexToAddress("0x4d3d5c850dd5bd9d6f4adda3dd039a3c8054ca29"),
		CoordinatorRegistry: common.HexToAddress("0xaa86dda78e9434aca114b6676fc742a18d15a1cc"),
		DevUtils:            common.HexToAddress("0xb23672f74749bf7916ba6827c64111a4d6de7f11"),
		WETH9:               common.HexToAddress("0x0b1ba0af832d7c05fd64161e0db78e85978e8082"),
		ZRXToken:            common.HexToAddress("0x871dd7c2b4b25e1aa18728e9d5f2af4c4e431f5c"),
		ChaiBridge:          common.HexToAddress("0x0000000000000000000000000000000000000000"),
		ChaiToken:           common.HexToAddress("0x0000000000000000000000000000000000000000"),
		MaximumGasPrice:     common.HexToAddress("0x2c530e4ecc573f11bd72cf5fdf580d134d25f15f"),
	}
}
