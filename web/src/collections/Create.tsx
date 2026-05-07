import { useState } from "react";
import AddIcon from "@mui/icons-material/Add";

import type { CreateCollectionRequest } from "../api/types";

import { useCollectionStore } from "../store/collection";

import Modal from "../components/ui/Modal";
import CircleButton from "../components/ui/CircleButton";

import styles from "./Create.module.css";
import Form from "../components/ui/Form";

const Create = () => {
  // API state wrapping
  const createCollection = useCollectionStore((s) => s.createCollection);

  const [modalOpen, setModalOpen] = useState(false);
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
    setModalOpen(false);
    setCreateCollectionData(null);
  };

  return (
    <>
      <CircleButton
        onClick={() => setModalOpen(true)}
        // text="Create Collection"
        icon={AddIcon}
      />

      {modalOpen && (
        <Modal className={styles.Modal} close={() => setModalOpen(false)}>
          <Form onSubmit={handleSubmit}>
            <Form.Field label="Name">
              <Form.Input
                id="collection-name"
                type="text"
                value={createCollectionData?.name}
                onChange={handleNameChange}
                placeholder="Enter Name"
                required
              />
            </Form.Field>
            <Form.Actions>
              <Form.Button type="submit">Create</Form.Button>
              <Form.Button type="button" onClick={() => setModalOpen(false)}>
                Cancel
              </Form.Button>
            </Form.Actions>
          </Form>
        </Modal>
      )}
    </>
  );
};

export default Create;
