import { ChangeEvent } from 'react';

export interface SubjectRowProps {
  fieldName: string;
  price: string;
  id: string;
  rowIndex: number;
  handleDelete: (rowId: string) => void;
  handleFieldChange: (
    rowId: string,
    fieldName: 'fieldName' | 'price'
  ) => (e: ChangeEvent<HTMLInputElement>) => void;
}

const SubjectRow = (props: SubjectRowProps) => {
  const { fieldName, price, handleDelete, rowIndex, handleFieldChange, id } =
    props;

  return (
    <div>
      <span>{rowIndex}</span>
      <input
        type="text"
        value={fieldName}
        onChange={(e) => {
          e.preventDefault();
          e.stopPropagation();

          handleFieldChange(id, 'fieldName')(e);
        }}
      />
      <input
        type="text"
        value={price}
        onChange={handleFieldChange(id, 'price')}
      />
      <button onClick={() => handleDelete(id)}>{'(x)'}</button>
    </div>
  );
};

export default SubjectRow;
