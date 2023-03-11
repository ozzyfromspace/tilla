import { StudentId } from '@/app/page';
import Button from '@/components/atoms/Button/Button';
import { SubjectRowProps } from '@/components/atoms/SubjectRow/SubjectRow';
import { BASE_URL } from '@/constants';
import { Combobox } from '@headlessui/react';
import axios from 'axios';
import { Dispatch, MouseEvent, SetStateAction, useState } from 'react';
import { v4 as uuidv4 } from 'uuid';
import SubjectRowGrid from '../SubjectRowGrid/SubjectRowGrid';

function toFullName(studentId: StudentId): string {
  return `${studentId.firstName} ${studentId.lastName}`;
}

const handleSetSelectedPersonId =
  (
    studentIds: StudentId[],
    setSelectedPersonId: Dispatch<SetStateAction<StudentId>>
  ) =>
  (fieldName: string) => {
    for (let i = 0; i < studentIds.length; i++) {
      console.log('in loop', i, studentIds, studentIds[0]);
      const studentName = toFullName(studentIds[i]);

      if (studentName === fieldName) {
        setSelectedPersonId(() => studentIds[i]);
        return;
      }
    }
  };

const handleCreateSubjectRow =
  (setSubjectRows: Dispatch<SetStateAction<SubjectRowProps[]>>) =>
  (e: MouseEvent<HTMLButtonElement, globalThis.MouseEvent>) => {
    e.preventDefault();
    e.stopPropagation();

    const newId = uuidv4();
    const getNewRowProps = (rowIndex: number): SubjectRowProps => ({
      fieldName: '',
      price: '',
      rowIndex: rowIndex,
      id: newId,
      handleDelete: () => {},
      handleFieldChange: () => () => {},
    });

    setSubjectRows((subs) => {
      return [...subs, getNewRowProps(subs.length + 1)];
    });
  };

const submitSubjects =
  (
    subjectRows: SubjectRowProps[],
    setSubjectRows: Dispatch<SetStateAction<SubjectRowProps[]>>,
    selectedPersonId: StudentId
  ) =>
  (e: MouseEvent<HTMLFormElement, globalThis.MouseEvent>) => {
    interface Subject {
      name: string;
      pricePerHour: number;
    }

    interface Payload {
      studentId: string;
      subjects: Subject[];
    }

    e.preventDefault();
    if (!selectedPersonId.id) {
      console.log('no student id');
      return;
    } else {
      console.log('made it!', selectedPersonId);
    }

    function coerse(si: SubjectRowProps[]): Subject[] {
      const subs: Subject[] = [];

      for (const row of si) {
        if (!row.fieldName || !row.price) continue;

        const sub: Subject = {
          name: row.fieldName,
          pricePerHour: parseFloat(parseFloat(row.price).toFixed(2)),
        };

        subs.push(sub);
      }

      return subs;
    }

    const payload: Payload = {
      studentId: selectedPersonId.id,
      subjects: coerse(subjectRows),
    };

    if (payload.subjects.length === 0) {
      console.log('no valid subject entries');
      return;
    }

    axios.post(`${BASE_URL}/student/subjects`, payload).then((resp) => {
      if (resp.status === 201) {
        setSubjectRows(() => []);
      }
    });
  };

interface Props {
  subjectRows: SubjectRowProps[];
  setSubjectRows: Dispatch<SetStateAction<SubjectRowProps[]>>;
  studentIds: StudentId[];
  selectedPersonId: StudentId;
  setSelectedPersonId: Dispatch<SetStateAction<StudentId>>;
}

const AddSubjectForm = (props: Props) => {
  const {
    studentIds,
    subjectRows,
    setSubjectRows,
    selectedPersonId,
    setSelectedPersonId,
  } = props;
  const [query, setQuery] = useState('');

  const filteredPeople =
    query === ''
      ? studentIds
      : studentIds.filter((studentId) => {
          return toFullName(studentId)
            .toLowerCase()
            .includes(query.toLowerCase());
        });

  return (
    <form
      onSubmit={submitSubjects(subjectRows, setSubjectRows, selectedPersonId)}
    >
      <Combobox
        value={toFullName(selectedPersonId)}
        onChange={handleSetSelectedPersonId(studentIds, setSelectedPersonId)}
      >
        <Combobox.Input onChange={(event) => setQuery(event.target.value)} />
        <Combobox.Options>
          {filteredPeople.map((person) => (
            <Combobox.Option key={person.id} value={toFullName(person)}>
              {toFullName(person)}
            </Combobox.Option>
          ))}
        </Combobox.Options>
      </Combobox>
      <Button
        label="New Subject"
        type="button"
        onClick={handleCreateSubjectRow(setSubjectRows)}
      />
      <SubjectRowGrid
        subjectRows={subjectRows}
        setSubjectRows={setSubjectRows}
      />
      <Button
        label="Save"
        type="submit"
        onClick={() => console.log(subjectRows, selectedPersonId)}
      />
    </form>
  );
};

export default AddSubjectForm;
