import { Teacher, TeacherField } from '@/app/page';
import Button from '@/components/atoms/Button/Button';
import RowInput from '@/components/atoms/RowInput/RowInput';
import { BASE_URL } from '@/constants';
import axios from 'axios';
import { ChangeEvent, Dispatch, FormEvent, SetStateAction } from 'react';

interface Props {
  teacher: Teacher;
  setTeacher: Dispatch<SetStateAction<Teacher>>;
}

const handleAddTeacherSubmit =
  (teacher: Teacher, setTeacher: Dispatch<SetStateAction<Teacher>>) =>
  (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!teacher.firstName) {
      console.log('no teacher firstname was provided');
      return;
    }

    if (!teacher.lastName) {
      console.log('no teacher lastname was provided');
      return;
    }

    if (!teacher.nickname) {
      console.log('no teacher nickname was provided');
      return;
    }

    axios.post(`${BASE_URL}/teacher`, teacher).then((resp) => {
      console.log(resp);

      if (resp.status === 201) {
        setTeacher(() => ({
          firstName: '',
          lastName: '',
          nickname: '',
        }));
      }
    });
  };

const handleTeacherChange = (
  inputField: TeacherField,
  setTeacher: Dispatch<SetStateAction<Teacher>>
) => {
  return (e: ChangeEvent<HTMLInputElement>) => {
    setTeacher((t) => ({ ...t, [inputField]: e.target.value }));
  };
};

const AddTeacherForm = (props: Props) => {
  const { teacher, setTeacher } = props;

  return (
    <form onSubmit={handleAddTeacherSubmit(teacher, setTeacher)}>
      <RowInput
        label="First Name"
        inputValue={teacher.firstName}
        handleChange={handleTeacherChange('firstName', setTeacher)}
      />
      <RowInput
        label="Last Name"
        inputValue={teacher.lastName}
        handleChange={handleTeacherChange('lastName', setTeacher)}
      />
      <RowInput
        label="Nickname"
        inputValue={teacher.nickname}
        handleChange={handleTeacherChange('nickname', setTeacher)}
      />
      <Button label="Add Teacher" />
    </form>
  );
};

export default AddTeacherForm;
