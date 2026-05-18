import { useState } from "react";
import { useParams } from "react-router";

import type { CreateItemRequest } from "../api/types";

import { useModalStore } from "../store/modal";
import { useItemStore } from "../store/items";

import Modal from "../components/ui/Modal";
import Form from "../components/ui/Form";

import styles from "./Items.module.css";

const CreateModal = () => {
  const { id } = useParams();

  // API state wrapping
  const createItem = useItemStore((s) => s.createItem);

  const [createItemData, setCreateItemData] =
    useState<CreateItemRequest | null>(null);

  const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setCreateItemData((prevState) => ({
      ...prevState,
      name: e.target.value,
    }));
  };

  const handleSubmit = (e: React.SubmitEvent<HTMLFormElement>) => {
    // Stop page refresh
    e.preventDefault();

    // Create the collection
    if (createItemData === null) return;
    createItem(id!, createItemData).catch(console.log);

    // Clean up after submit
    close();
    setCreateItemData(null);
  };

  const close = useModalStore((s) => s.closeModal);

  return (
    <Modal className={styles.Modal} close={close}>
      <Form onSubmit={handleSubmit}>
        <Form.Field label="Name">
          <Form.Input
            id="item-name"
            type="text"
            value={createItemData?.name ?? ""}
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
