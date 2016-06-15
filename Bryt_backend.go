package main

import (
	"fmt"
	"net/http"
	"github.com/eauge/opentok-go-sdk"
)

func activeSessions(request, response)  {

	/* need to see interface and/or implementation for request 
	and response */

	activeSession := request.object
	session_ID := activeSession.ArchiveGet("sessionID")
	if (session_ID) {
		response.success()
		return
	}

	ot.Session(func (err, sessionId) {
		err_m := "unable to create opentok session for: " + activeSession.ID
		if (err) {
			response.error(err_m)
			return
		}
		activeSession.SessionID := session_ID

		/* now generate token */

		publisherToken := ot.Token(session_ID, { "role" : ot.ROLE.PUBLISHER})
		/* what is the JSON object as a parameter for? */
		if (!publisherToken) {
			response.error("unable to create publisher token for: " + activeSession.ID)
			return
		}

		subscriberToken := ot.Token(session_ID, { "role" : ot.ROLE.SUBSCRIBER })
		/* why json object as param? */
		if (!subscriberToken) {
			response.error("unable to create subscriber token for: " + activeSession.ID)
			return
		}

		/* save session_ID in the activeSessions object */

		activeSession.publisherToken := publisherToken
		activeSession.subscriberToken := subscriberToken
		response.success()
	})
})


func getActiveSessionsToken (request, response) {
	activeSessionId := request.params.activeSession 
	//what type is request and response?
	if (!activeSessionId) { 
		response.error("must provide an activeSession object id")
	}
	// activeSessionQuery := new Parse.Query("ActiveSessions")

	/*activeSessionQuery.get(....
	
		more code using activeSessionQuery 
		but activeSessionQuery is a query connected 
		to the Parse methods/objects
		so irrevelant in this Go backend file
	)
	*/


}

//helper func to figure out the opentok role a user
// should get based on the Boradcast object
roleForUser := func (activeSession, user) {
	// an activeSessions owner gets a publisher token
	if (activeSession.callerID == user.id) { 
		return ot.ROLE.PUBLISHER
	}
	// else give Subscriber token
	else {
		return ot.ROLE.SUBSCRIBER
	}
}

func main {
	apiKey := 45490452
	apiSecret := "a816ffe35757b76b2287c24c567e4b6803351647"
	ot := opentok.New(apiKey, apiSecret)

	s, err := ot.Session(nil)
	if err != nil {
    	panic(err)
	}

	t, err := ot.Token(s, nil)
	if err != nil {
    	panic(err)
	}

	fmt.Println("session: ", s.ID)
	fmt.Println("token: ", t)
}