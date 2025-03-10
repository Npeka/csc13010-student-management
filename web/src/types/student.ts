export interface Student {
  student_id: string;
  full_name: string;
  birth_date: string;
  gender: string;
  faculty: string;
  course: string;
  program: number;
  address: string;
  email: string;
  phone: string;
  status: number;
}

export interface StudentResponseDTO {
  id: number;
  student_id: string;
  full_name: string;
  birth_date: string;
  gender_id: number;
  faculty_id: number;
  course_id: number;
  program_id: number;
  address: string;
  email: string;
  phone: string;
  status_id: number;
  created_at: string;
  updated_at: string;

  gender?: string;
  faculty?: string;
  course?: string;
  program?: string;
  status?: string;
}

export interface Option {
  id: number;
  name: string;
}

export interface OptionDTO {
  genders: Option[];
  faculties: Option[];
  courses: Option[];
  programs: Option[];
  statuses: Option[];
}

export interface Faculty {
  id?: string;
  name: string;
}

export interface Program {
  id?: string;
  name: string;
}

export interface Status {
  id?: string;
  name: string;
}
