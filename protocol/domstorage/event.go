// Code generated by cdpgen. DO NOT EDIT.

package domstorage

import (
	"github.com/mafredri/cdp/rpcc"
)

// ItemsClearedClient is a client for DOMStorageItemsCleared events.
type ItemsClearedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ItemsClearedReply, error)
	rpcc.Stream
}

// ItemsClearedReply is the reply for DOMStorageItemsCleared events.
type ItemsClearedReply struct {
	StorageID StorageID `json:"storageId"` // No description.
}

// ItemRemovedClient is a client for DOMStorageItemRemoved events.
type ItemRemovedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ItemRemovedReply, error)
	rpcc.Stream
}

// ItemRemovedReply is the reply for DOMStorageItemRemoved events.
type ItemRemovedReply struct {
	StorageID StorageID `json:"storageId"` // No description.
	Key       string    `json:"key"`       // No description.
}

// ItemAddedClient is a client for DOMStorageItemAdded events.
type ItemAddedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ItemAddedReply, error)
	rpcc.Stream
}

// ItemAddedReply is the reply for DOMStorageItemAdded events.
type ItemAddedReply struct {
	StorageID StorageID `json:"storageId"` // No description.
	Key       string    `json:"key"`       // No description.
	NewValue  string    `json:"newValue"`  // No description.
}

// ItemUpdatedClient is a client for DOMStorageItemUpdated events.
type ItemUpdatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ItemUpdatedReply, error)
	rpcc.Stream
}

// ItemUpdatedReply is the reply for DOMStorageItemUpdated events.
type ItemUpdatedReply struct {
	StorageID StorageID `json:"storageId"` // No description.
	Key       string    `json:"key"`       // No description.
	OldValue  string    `json:"oldValue"`  // No description.
	NewValue  string    `json:"newValue"`  // No description.
}
