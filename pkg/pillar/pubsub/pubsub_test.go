package pubsub

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Item struct {
	aString string
}

var item = Item{
	aString: "aString",
}
var sub, err = Subscribe("agentName", item, false, &item)

func TestHandleModify(t *testing.T) {
	sub.agentScope = "agentScope"
	sub.topic = "topic"
	created := false
	modified := false
	subCreateHandler := func(ctxArg interface{}, key string, status interface{}) {
		created = true
	}
	subModifyHandler := func(ctxArg interface{}, key string, status interface{}) {
		modified = true
	}

	testMatrix := map[string]struct {
		ctxArg         Subscription
		key            string
		item           interface{}
		modifyHandler  SubHandler
		createHandler  SubHandler
		expectedCreate bool
		expectedModify bool
	}{
		"Modify Handler is nil": {
			ctxArg:         *sub,
			key:            "key_0",
			item:           item,
			modifyHandler:  nil,
			createHandler:  subCreateHandler,
			expectedCreate: true,
			expectedModify: false,
		},
		"Create Handler is nil": {
			ctxArg:         *sub,
			key:            "key_1",
			item:           item,
			modifyHandler:  subModifyHandler,
			createHandler:  nil,
			expectedCreate: false,
			expectedModify: true,
		},
		"Create Handler and Modify Handler are nil": {
			ctxArg:         *sub,
			key:            "key_2",
			item:           item,
			modifyHandler:  nil,
			createHandler:  nil,
			expectedCreate: false,
			expectedModify: false,
		},
	}
	for testname, test := range testMatrix {
		t.Logf("Running test case %s", testname)
		test.ctxArg.CreateHandler = test.createHandler
		test.ctxArg.ModifyHandler = test.modifyHandler
		handleModify(&test.ctxArg, test.key, test.item)
		// Make sure both weren't called
		assert.Equal(t, created, test.expectedCreate)
		assert.Equal(t, modified, test.expectedModify)
		// Reset created and modified to false for next test
		created = false
		modified = false
	}
}
