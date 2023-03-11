import Button from '@/components/atoms/Button/Button';
import { SubjectRowProps } from '@/components/atoms/SubjectRow/SubjectRow';
import { BASE_URL } from '@/constants';
import { StudentId, useStudents } from '@/hooks';
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
    studentDataSlice: StudentId[],
    setSelectedStudent: Dispatch<SetStateAction<StudentId>>
  ) =>
  (fieldName: string) => {
    for (let i = 0; i < studentDataSlice.length; i++) {
      console.log('in loop', i, studentDataSlice, studentDataSlice[0]);
      const studentName = toFullName(studentDataSlice[i]);

      if (studentName === fieldName) {
        setSelectedStudent(() => studentDataSlice[i]);
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
    selectedStudent: StudentId
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
    if (!selectedStudent.id) {
      console.log('no student id');
      return;
    } else {
      console.log('made it!', selectedStudent);
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
      studentId: selectedStudent.id,
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

const AddSubjectForm = () => {
  const { selectedStudent, setSelectedStudent, studentDataSlice } =
    useStudents();

  const [query, setQuery] = useState('');
  const [subjectRows, setSubjectRows] = useState<SubjectRowProps[]>(() => []);

  const filteredPeople =
    query === ''
      ? studentDataSlice
      : studentDataSlice.filter((studentId) => {
          return toFullName(studentId)
            .toLowerCase()
            .includes(query.toLowerCase());
        });

  return (
    <form
      onSubmit={submitSubjects(subjectRows, setSubjectRows, selectedStudent)}
    >
      <div className="flex justify-between gap-3 p-3 bg-slate-100 max-w-lg rounded-md">
        <Combobox
          value={toFullName(selectedStudent)}
          onChange={handleSetSelectedPersonId(
            studentDataSlice,
            setSelectedStudent
          )}
        >
          <Combobox.Input
            onChange={(event) => setQuery(event.target.value)}
            className="flex-1"
            placeholder="Select Student"
          />
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
          selected={false}
        />
        <Button
          label="Save"
          type="submit"
          onClick={() => console.log(subjectRows, selectedStudent)}
          selected
        />
      </div>
      <SubjectRowGrid
        subjectRows={subjectRows}
        setSubjectRows={setSubjectRows}
      />
    </form>
  );
};

export default AddSubjectForm;
