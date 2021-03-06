// Code generated by "callbackgen -type Telegram"; DO NOT EDIT.

package interact

import ()

func (tm *Telegram) OnAuthorized(cb func(a *TelegramAuthorizer)) {
	tm.authorizedCallbacks = append(tm.authorizedCallbacks, cb)
}

func (tm *Telegram) EmitAuthorized(a *TelegramAuthorizer) {
	for _, cb := range tm.authorizedCallbacks {
		cb(a)
	}
}
