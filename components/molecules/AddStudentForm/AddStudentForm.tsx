import Button from '@/components/atoms/Button/Button';
import RowInput from '@/components/atoms/RowInput/RowInput';
import { BASE_URL } from '@/constants';
import { useStudents } from '@/hooks';
import axios from 'axios';
import {
  ChangeEvent,
  Dispatch,
  FormEvent,
  SetStateAction,
  useState,
} from 'react';

export interface Student {
  firstName: string;
  lastName: string;
  nickname: string;
  calendarId: string;
}

export type StudentField = keyof Student;

const handleAddStudentSubmit =
  (
    student: Student,
    setStudent: Dispatch<SetStateAction<Student>>,
    fetchStudents: () => Promise<void>
  ) =>
  (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!student.firstName) {
      console.log('no student firstname was provided');
      return;
    }

    if (!student.lastName) {
      console.log('no student lastname was provided');
      return;
    }

    if (!student.nickname) {
      console.log('no student nickname was provided');
      return;
    }

    if (!student.calendarId) {
      console.log('no student calendar id was provided');
      return;
    }

    axios.post(`${BASE_URL}/student`, student).then((resp) => {
      console.log(resp);

      if (resp.status === 201) {
        fetchStudents();

        setStudent(() => ({
          firstName: '',
          lastName: '',
          nickname: '',
          calendarId: '',
        }));
      }
    });
  };

const handleStudentChange = (
  inputField: StudentField,
  setStudent: Dispatch<SetStateAction<Student>>
) => {
  return (e: ChangeEvent<HTMLInputElement>) => {
    setStudent((s) => ({ ...s, [inputField]: e.target.value }));
  };
};

const AddStudentForm = () => {
  const { fetchStudents } = useStudents();
  const [student, setStudent] = useState<Student>(() => ({
    firstName: '',
    lastName: '',
    nickname: '',
    calendarId: '',
  }));

  return (
    <form
      onSubmit={handleAddStudentSubmit(student, setStudent, fetchStudents)}
      className="space-y-2"
    >
      <RowInput
        label="First Name"
        inputValue={student.firstName}
        handleChange={handleStudentChange('firstName', setStudent)}
        placeholder="Toby"
      />
      <RowInput
        label="Last Name"
        inputValue={student.lastName}
        handleChange={handleStudentChange('lastName', setStudent)}
        placeholder="Maguire"
      />
      <RowInput
        label="Nickname"
        inputValue={student.nickname}
        handleChange={handleStudentChange('nickname', setStudent)}
        placeholder="spidey"
      />
      <RowInput
        label="Calendar Id"
        inputValue={student.calendarId}
        handleChange={handleStudentChange('calendarId', setStudent)}
        placeholder="id@marvel.com"
      />
      <div className="flex justify-center items-center pt-6">
        <Button label="Add Student" />
      </div>
    </form>
  );
};

export default AddStudentForm;
