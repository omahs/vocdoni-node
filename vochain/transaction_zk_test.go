package vochain

import (
	"crypto/sha256"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/arbo"
	snarkParsers "github.com/vocdoni/go-snark/parsers"
	"github.com/vocdoni/go-snark/types"
	"go.vocdoni.io/dvote/vochain/transaction/vochaintx"
	models "go.vocdoni.io/proto/build/go/models"
)

func TestVoteCheckZkSNARK(t *testing.T) {
	app := TestBaseApplication(t)

	vkJSON := `
	{
		"protocol": "groth16",
		"curve": "bn128",
		"nPublic": 7,
		"vk_alpha_1": [
		 "16690211343505291795059441371620861229164781336230094380457367531934880398502",
		 "18687425750113507431307649431006465296160825402520895668098920889979328769547",
		 "1"
		],
		"vk_beta_2": [
		 [
		  "14207705963228831532911208547908707881857764048673802573407997316232016069645",
		  "5169946860532989220962328004639764432258742190027053377785212905803865274740"
		 ],
		 [
		  "4197538776290921693016628076984378284386107112507027312002666476718723315728",
		  "14982486346550771544786599717754917898007168315267333763571937384967020148718"
		 ],
		 [
		  "1",
		  "0"
		 ]
		],
		"vk_gamma_2": [
		 [
		  "10857046999023057135944570762232829481370756359578518086990519993285655852781",
		  "11559732032986387107991004021392285783925812861821192530917403151452391805634"
		 ],
		 [
		  "8495653923123431417604973247489272438418190587263600148770280649306958101930",
		  "4082367875863433681332203403145435568316851327593401208105741076214120093531"
		 ],
		 [
		  "1",
		  "0"
		 ]
		],
		"vk_delta_2": [
		 [
		  "14515656650793456285574892197079004502820008859644458327663450631995165810480",
		  "17813943130201805344141732176669074691912364024901444006322060883610240967201"
		 ],
		 [
		  "1516161148717103475120421544631641914510838488164696481159607190066366046884",
		  "11484953488255333787814603624041635168806752101863078754223759660795980715618"
		 ],
		 [
		  "1",
		  "0"
		 ]
		],
		"vk_alphabeta_12": [
		 [
		  [
		   "13637451605841267050718292143833234882903048452950890439594814037598694655073",
		   "11014379352423758232580740412230823089396385193947477602886458004186447426362"
		  ],
		  [
		   "11764131589963512325185350828969172058139776873471186918550949437635493180378",
		   "9537806181251743766221625187339749057927266008332183723175168854140667873013"
		  ],
		  [
		   "12323387886996089659318315423189827935759337676800395401518555428630707159065",
		   "3755519908163389474707985652515238243296292900454238378610233994623907177720"
		  ]
		 ],
		 [
		  [
		   "397806229012974097484422551939718648405545801750939874817162728931996802132",
		   "360986825744150890257816492588070101912038786039570693970855130512326721546"
		  ],
		  [
		   "20569965811872242697745192038255134541810751222115475944631106624903990534095",
		   "11989503282161871621731994055257983231489826319186702024264453234578107893100"
		  ],
		  [
		   "20476548451897298370931827127997719116240177997601934258364937517638868349525",
		   "14975453099091672574099554417240015193432336025628515180121236038046280291629"
		  ]
		 ]
		],
		"IC": [
		 [
		  "15922585095232984610643076050471875533912047901682345213797824762718899501138",
		  "20448316784731168508980979076065578875725352385623729046376018460249108618291",
		  "1"
		 ],
		 [
		  "4088689542895364999208932514094652937151046481083961838536935829232191514435",
		  "15392351696316395842352353541090500651418669429244406461864999807079401455319",
		  "1"
		 ],
		 [
		  "19319516116066201815303141942306031683526865656914699514509377803394798142666",
		  "10543563336996341814983066693731700315512544706200505628528990795505144137478",
		  "1"
		 ],
		 [
		  "6449705476868668114303941230166629516657533015458256645971380207217165712746",
		  "8740649991356532701110453713660735019886802930551050466279603687411116603140",
		  "1"
		 ],
		 [
		  "6971474991519439609075821107602485917699763040449303156418543609925079131292",
		  "21418256113359646273640910437579058854380272553637003915892188330367809183311",
		  "1"
		 ],
		 [
		  "21135216502371639305390890818504991238339780767112560788652058660747977117027",
		  "20156746729173486464427451903235411242380606131979858681394358960416280622235",
		  "1"
		 ],
		 [
		  "11055322257434982384270182163886277319913062583176596166783023895934016271287",
		  "12750302130334196541559902248458899673899311048479947513387215266188850657312",
		  "1"
		 ],
		 [
		  "7206074413591916461066451896031306378443051806701563800629002084402821473823",
		  "922170089553272762260291976229261520758656683175931138765611212548134612128",
		  "1"
		 ]
		]
	   }
	`
	vk0, err := snarkParsers.ParseVk([]byte(vkJSON))
	qt.Assert(t, err, qt.IsNil)
	app.TransactionHandler.ZkVKs = []*types.Vk{vk0}

	processId := sha256.Sum256(big.NewInt(10).Bytes()) // processId is a byte-array of 32 bytes
	entityId := []byte("entityid-test")
	censusRootBI, ok := new(big.Int).SetString("10880001835876045062663123179921953501658252221549048512484733871519188473352", 10)
	qt.Assert(t, ok, qt.IsTrue)

	process := &models.Process{
		ProcessId: processId[:],
		EntityId:  entityId,
		EnvelopeType: &models.EnvelopeType{
			Anonymous: true,
		},
		Mode:       &models.ProcessMode{},
		Status:     models.ProcessStatus_READY,
		CensusRoot: make([]byte, 32), // emtpy hash
		StartBlock: 1,
		BlockCount: 3,
	}
	err = app.State.AddProcess(process)
	qt.Assert(t, err, qt.IsNil)
	process, err = app.State.Process(processId[:], false)
	qt.Assert(t, err, qt.IsNil)
	process.RollingCensusRoot = arbo.BigIntToBytes(32, censusRootBI)
	err = app.State.UpdateProcess(process, processId[:])
	qt.Assert(t, err, qt.IsNil)

	// proof data generated from js (snarkjs)
	protoProof := models.ProofZkSNARK{
		CircuitParametersIndex: 0,
		A: []string{
			"3760363242862841938281052770301817413473662426206426210307710005254968812065",
			"13258033588191402629745415658105736164193886576672483600930357465360082629644",
			"1",
		},
		B: []string{
			"6726472988384325527312583364100254522475982615851595557547287705425718401220",
			"11903517080030121863728171464173531510900006478042346238807022130328955049926",
			"12847079423552046877419374461987024132336480533520565301368818108517092267359",
			"7506410150058374479092761256683660342325877506195081689394850762462031946066",
			"1",
			"0",
		},
		C: []string{
			"20803962252196751695731425118836885562625495317417574516468043201927542167246",
			"18868334046974605488605481850593847767824122322911317209649976920575418827992",
			"1",
		},
		PublicInputs: []string{
			"242108076058607163538102198631955675649",
			"142667662805314151155817304537028292174",
			"10880001835876045062663123179921953501658252221549048512484733871519188473352",
			"4295509861249984880361571032347194270863089509149412623993065795982837479793",
			"1",
			"302689215824177652345211539748426020171",
			"205062086841587857568430695525160476881",
		},
	}

	nullifierBI, ok := new(big.Int).SetString("4295509861249984880361571032347194270863089509149412623993065795982837479793", 10)
	qt.Assert(t, ok, qt.IsTrue)
	nullifier := arbo.BigIntToBytes(32, nullifierBI)

	voteValue := big.NewInt(1).Bytes()
	vtx := &models.VoteEnvelope{
		ProcessId:   processId[:],
		VotePackage: voteValue,
		Nullifier:   nullifier,
		Proof: &models.Proof{
			Payload: &models.Proof_ZkSnark{
				ZkSnark: &protoProof,
			},
		},
	}
	signature := []byte{}
	txBytes := []byte{}
	txID := [32]byte{}
	commit := false

	_, _, err = app.TransactionHandler.VoteTxCheck(&vochaintx.VochainTx{
		Tx:         &models.Tx{Payload: &models.Tx_Vote{Vote: vtx}},
		Signature:  signature,
		SignedBody: txBytes,
		TxID:       txID,
	}, commit)
	qt.Assert(t, err, qt.IsNil)
}
