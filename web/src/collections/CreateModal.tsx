import { useState } from "react";
import { useNavigate } from "react-router";

import type { CreateCollectionRequest } from "../api/types";

import { useCollectionStore } from "../store/collection";

import Modal from "../components/ui/Modal";
import Form from "../components/ui/Form";

import styles from "./Collections.module.css";

const CreateModal = () => {
  const navigate = useNavigate();

  // API state wrapping
  const createCollection = useCollectionStore((s) => s.createCollection);

  const [createCollectionData, setCreateCollectionData] =
    useState<CreateCollectionRequest | null>(null);

  const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setCreateCollectionData((prevState) => ({
      ...prevState,
      name: e.target.value,
    }));
  };

  const handleSubmit = (e: React.SubmitEvent<HTMLFormElement>) => {
    // Stop page refresh
    e.preventDefault();

    // Create the collection
    if (createCollectionData === null) return;
    createCollection(createCollectionData).catch(console.log);

    // Clean up after submit
    close();
    setCreateCollectionData(null);
  };

  const close = () => navigate("..");

  return (
    <Modal className={styles.Modal} close={close}>
      <Form onSubmit={handleSubmit}>
        <Form.Field label="Name">
          <Form.Input
            id="collection-name"
            type="text"
            value={createCollectionData?.name ?? ""}
            onChange={handleNameChange}
            placeholder="Enter Name"
            required
          />
        </Form.Field>
        <Form.Actions>
          <Form.Button type="submit">Create</Form.Button>
          <Form.Button type="button" onClick={close}>
            Cancel
          </Form.Button>
        </Form.Actions>
      </Form>
    </Modal>
  );
};

export default CreateModal;
