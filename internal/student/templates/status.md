# GIẤY XÁC NHẬN TÌNH TRẠNG SINH VIÊN

**Trường:** {{ .University }}  
**Địa chỉ:** {{ .Address }}  
**Điện thoại:** {{ .Phone }} | **Email:** {{ .Email }}

## 1. Thông tin sinh viên:

- **Họ và tên:** {{ .Student.FullName }}
- **Mã số sinh viên:** {{ .Student.StudentID }}
- **Ngày sinh:** {{ .Student.BirthDate }}
- **Giới tính:** {{ .Student.Gender }}
- **Khoa:** {{ .Student.Faculty }}
- **Chương trình:** {{ .Student.Program }}
- **Tình trạng:** {{ .Student.Status }}

## 2. Xác nhận

- Giấy xác nhận có hiệu lực đến ngày: {{ .ValidUntil }}

---

**Trưởng phòng đào tạo**  
_(Ký tên, đóng dấu)_
