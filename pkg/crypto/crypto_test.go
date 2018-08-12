package crypto

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

const secretPhrase = "glad suffer red during single glow shut slam hill death lust although"

func hexToBytes(s string) []byte {
	bs, _ := hex.DecodeString(s)
	return bs
}

func TestSecretPhraseToPrivateKey(t *testing.T) {
	privateKey := secretPhraseToPrivateKey(secretPhrase)
	assert.Equal(t, "e04c16cfd3d1cbf11a51fa0c75c09d43d307a10ae5149b1ec0ddba661b9d2f5e",
		hex.EncodeToString(privateKey[:]))
}

func TestSecretPhraseToPublicKey(t *testing.T) {
	publicKey := secretPhraseToPublicKey(secretPhrase)
	assert.Equal(t, "a9fc9b42e3918c3e109fdc819aa092fd06ce7c18d653bf003b31d48d35699104",
		hex.EncodeToString(publicKey[:]))
}

func TestSign(t *testing.T) {
	msg := make([]byte, 128)
	for i := 0; i < 128; i++ {
		msg[i] = byte(i)
	}
	sig := sign(msg, secretPhrase)
	assert.Equal(t, "cea74dd7177ddd8bd88e4a0d8807da533e95195d97cc4bcb0f946f355a8c85035ebe57fe9628f93405cf9f7e83bfc64968bc992d3508434a3023fba6b6cab491", (hex.EncodeToString(sig)))
}

func TestVerify(t *testing.T) {
	type testCase struct {
		sig    string
		pubKey string
		msg    string
	}

	tcs := []*testCase{
		&testCase{
			sig:    "8B56C7E025692298118B46B13934DC913736A82B7E03926CEA3E94F9C2F939036D2FB38E5AAFC4904A25B60DFCCFDE58EA5C044DB03FB652D538D87F37D594A7",
			pubKey: "8C22857E71B401212FC95A340198C61A11CFA17F6553F5101BD2EFF49981DB4B",
			msg:    "00002EC20000A0058C22857E71B401212FC95A340198C61A11CFA17F6553F5101BD2EFF49981DB4B20397FC3C03E152400460A99E800000000CA9A3B00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		},
		&testCase{
			sig:    "F1E2128A8E2CFA099407B0595CCE49C03AE89F40B731ADEFEBC38FF023351A0D8F9E3B52F3B68DD2F552E209ADCC829F70F9E58B48F5585C56FE2D57355D8B0A",
			pubKey: "18BC3D3E3B7AD01A77FB85315330397413C0052FC403CB2C555C87EFED2B897E",
			msg:    "00001EE50000A00518BC3D3E3B7AD01A77FB85315330397413C0052FC403CB2C555C87EFED2B897E370D00E5722B437400E1F5050000000000E1F50500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		},
		&testCase{
			sig:    "CDA595EA0AD2739C7DBFBBB9A01067BC757AA288ECB6EB60938B172E6466220212ADA84F380DB1C76D69D7EFDF211E7835274D7A76ADFD8189D26DF041E62AF6",
			pubKey: "18BC3D3E3B7AD01A77FB85315330397413C0052FC403CB2C555C87EFED2B897E",
			msg:    "000077E50000A00518BC3D3E3B7AD01A77FB85315330397413C0052FC403CB2C555C87EFED2B897E1B812086BEDFF26A00A3E1110000000000E1F50500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		},
		&testCase{
			sig:    "D14B7671B1539E16F6FE975D9DFDDB7D48B08C5EA28A96E5E27B145188A943035A1FF64381D1B4E406485E5748D0DA34A56D6D8A9B03AEA125C68F88AD9DA985",
			pubKey: "18BC3D3E3B7AD01A77FB85315330397413C0052FC403CB2C555C87EFED2B897E",
			msg:    "000060E50000A00518BC3D3E3B7AD01A77FB85315330397413C0052FC403CB2C555C87EFED2B897EE7DD9DFAA8F664A700C2EB0B0000000000E1F50500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		},
		&testCase{
			sig:    "4444CAAF7F0996A1DCD37AEC29D07CD8BDC964159F175D2D25EE3F49D4C69504C353109AFE21F6FA28E01BC1B1A2279C270E4890B5A3AC7A887E2401B753D7F2",
			pubKey: "EBCE85E49EA725B214C909FD9969570EB8F456D804B22B25C2A0DBFD0D31EC4C",
			msg:    "010057F20000A005EBCE85E49EA725B214C909FD9969570EB8F456D804B22B25C2A0DBFD0D31EC4CD349D0629D29276F000000000000000000E1F5050000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000037000000596F752066696E64206D616E7920626C6F636B732C2063616E20796F752073656E64206D6520736F6D6520636F696E20706C6561736521",
		},
		&testCase{
			sig:    "737808F2DD170739E9D9CB6015889E22C3F159AEE29B77C98A4A22AF18091D0A37758FBDF1BBE5ACEAF38954C0247AF842C8EBEC7768E9C93B2A9B7CE198B9BA",
			pubKey: "EBCE85E49EA725B214C909FD9969570EB8F456D804B22B25C2A0DBFD0D31EC4C",
			msg:    "0100B7F20000A005EBCE85E49EA725B214C909FD9969570EB8F456D804B22B25C2A0DBFD0D31EC4CC9C550D65643EF77000000000000000000E1F5050000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000037000000596F752066696E64206D616E7920626C6F636B732C2063616E20796F752073656E64206D6520736F6D6520636F696E20706C6561736521",
		},
	}
	for _, tc := range tcs {
		sig := hexToBytes(tc.sig)
		pubKey := hexToBytes(tc.pubKey)
		msg := hexToBytes(tc.msg)
		assert.True(t, verify(sig, msg, pubKey, true))
	}
}

func BenchmarkSecretPhraseToPrivateKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		secretPhraseToPrivateKey(secretPhrase)
	}
}

func BenchmarkSecretPhraseToPublicKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		secretPhraseToPublicKey(secretPhrase)
	}
}

func BenchmarkSign(b *testing.B) {
	msg := make([]byte, 128)
	for i := 0; i < 128; i++ {
		msg[i] = byte(i)
	}
	for i := 0; i < b.N; i++ {
		sign(msg, secretPhrase)
		for j := 0; j < 128; j++ {
			msg[j] = byte(i + j)
		}
	}
}
