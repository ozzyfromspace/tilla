import { Student, StudentField } from '@/app/page';
import Button from '@/components/atoms/Button/Button';
import RowInput from '@/components/atoms/RowInput/RowInput';
import { BASE_URL } from '@/constants';
import axios from 'axios';
import { ChangeEvent, Dispatch, FormEvent, SetStateAction } from 'react';

interface Props {
  fetchStudents: () => Promise<void>;
  student: Student;
  setStudent: Dispatch<SetStateAction<Student>>;
}

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

const AddStudentForm = (props: Props) => {
  const { student, setStudent, fetchStudents } = props;

  return (
    <form onSubmit={handleAddStudentSubmit(student, setStudent, fetchStudents)}>
      <RowInput
        label="First Name"
        inputValue={student.firstName}
        handleChange={handleStudentChange('firstName', setStudent)}
      />
      <RowInput
        label="Last Name"
        inputValue={student.lastName}
        handleChange={handleStudentChange('lastName', setStudent)}
      />
      <RowInput
        label="Nickname"
        inputValue={student.nickname}
        handleChange={handleStudentChange('nickname', setStudent)}
      />
      <RowInput
        label="Calendar Id"
        inputValue={student.calendarId}
        handleChange={handleStudentChange('calendarId', setStudent)}
      />
      <Button label="Add Student" />
    </form>
  );
};

export default AddStudentForm;
