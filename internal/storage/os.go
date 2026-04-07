package storage

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"os"

// 	"github.com/rexgreenway/collection-app/internal/entities"
// )

// const OS StorageType = "os"

// // osStorage implements the Storage interface that reads & writes data from a
// // a JSON file.
// type osStorage struct {
// 	filePath string
// 	store    map[string]entities.Collection
// }

// // newOSStorage returns an implementation of the Storage interface that reads &
// // writes data from a file defined by the 'COLLECTION_STORAGE_FILE' environment
// // variable.
// func newOSStorage() (*osStorage, error) {
// 	// Read OS File Required Env Var
// 	filePathEnvVar := "COLLECTION_STORAGE_FILE"
// 	filePath := os.Getenv(filePathEnvVar)
// 	if filePath == "" {
// 		return &osStorage{}, fmt.Errorf("environment variable %q not set", filePathEnvVar)
// 	}

// 	// Create storage manager and read store
// 	storageManager := &osStorage{
// 		filePath: filePath,
// 	}
// 	if err := storageManager.ReadStore(); err != nil {
// 		return &osStorage{}, err
// 	}

// 	return storageManager, nil
// }

// // ReadStore reads the data from the instantiated osStorage file path file
// // path into the store.
// func (s *osStorage) ReadStore() error {
// 	f, err := os.Open(s.filePath)
// 	if err != nil {
// 		return err
// 	}
// 	encoder := json.NewDecoder(f)
// 	return encoder.Decode(&s.store)
// }

// // WriteStore writes the state of the current store to the file path
// // established at osStorage instantiation.
// func (s *osStorage) WriteStore() error {
// 	f, err := os.Create(s.filePath)
// 	if err != nil {
// 		return err
// 	}
// 	encoder := json.NewEncoder(f)
// 	encoder.SetIndent("", "    ")
// 	return encoder.Encode(s.store)
// }

// // ListCollections ???
// func (s *osStorage) ListCollections() (map[string]entities.Collection, error) {
// 	return s.store, nil
// }

// // CreateCollection ???
// func (s *osStorage) CreateCollection(collection entities.Collection) (entities.Collection, error) {
// 	if _, ok := s.store[collection.ID]; ok {
// 		return entities.Collection{}, ErrCollectionAlreadyExists
// 	}

// 	// Store & write
// 	s.store[collection.ID] = collection
// 	if err := s.WriteStore(); err != nil {
// 		return entities.Collection{}, err
// 	}

// 	fmt.Println(s.store)

// 	return collection, nil
// }

// // GetCollection ???
// func (s *osStorage) GetCollection(collectionID string) (entities.Collection, error) {
// 	collection, ok := s.store[collectionID]
// 	if !ok {
// 		return entities.Collection{}, errors.New("collection not found")
// 	}
// 	return collection, nil
// }
