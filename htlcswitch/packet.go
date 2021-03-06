package htlcswitch

import (
	"github.com/lightningnetwork/lnd/lnwire"
)

// htlcPacket is a wrapper around htlc lnwire update, which adds additional
// information which is needed by this package.
type htlcPacket struct {
	// destNode is the first-hop destination of a local created HTLC add
	// message.
	destNode [33]byte

	// incomingChanID is the ID of the channel that we have received an incoming
	// HTLC on.
	incomingChanID lnwire.ShortChannelID

	// outgoingChanID is the ID of the channel that we have offered or will
	// offer an outgoing HTLC on.
	outgoingChanID lnwire.ShortChannelID

	// incomingHTLCID is the ID of the HTLC that we have received from the peer
	// on the incoming channel.
	incomingHTLCID uint64

	// incomingHtlcAmt is the value of the *incoming* HTLC. This will be
	// set by the link when it receives an incoming HTLC to be forwarded
	// through the switch. Then the outgoing link will use this once it
	// creates a full circuit add. This allows us to properly populate the
	// forwarding event for this circuit/packet in the case the payment
	// circuit is successful.
	incomingHtlcAmt lnwire.MilliSatoshi

	// outgoingHTLCID is the ID of the HTLC that we offered to the peer on the
	// outgoing channel.
	outgoingHTLCID uint64

	// amount is the value of the HTLC that is being created or modified.
	amount lnwire.MilliSatoshi

	// htlc lnwire message type of which depends on switch request type.
	htlc lnwire.Message

	// obfuscator contains the necessary state to allow the switch to wrap
	// any forwarded errors in an additional layer of encryption.
	obfuscator ErrorEncrypter

	// localFailure is set to true if an HTLC fails for a local payment before
	// the first hop. In this case, the failure reason is simply encoded, not
	// encrypted with any shared secret.
	localFailure bool

	// isRouted is set to true if the incomingChanID and incomingHTLCID fields
	// of a forwarded fail packet are already set and do not need to be looked
	// up in the circuit map.
	isRouted bool

	// isResolution is set to true if this packet was actually an incoming
	// resolution message from an outside sub-system. We'll treat these as
	// if they emanated directly from the switch. As a result, we'll
	// encrypt all errors related to this packet as if we were the first
	// hop.
	isResolution bool
}
