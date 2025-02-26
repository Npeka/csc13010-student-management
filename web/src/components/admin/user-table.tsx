import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table";
import { Button } from "@/components/ui/button";
import { User } from "@/lib/model/user";
import { getUsers } from "@/lib/api/admin";
import { DeleteItem } from "@/components/admin/button";

export default async function UserTable() {
    const users = await getUsers();

    const columns = [
        { title: "Full Name", key: "full_name" },
        { title: "Email", key: "email" },
        { title: "Gender", key: "gender" },
        { title: "Date of Birth", key: "day_of_birth" },
        { title: "Created At", key: "created_at" },
        { title: "Updated At", key: "updated_at" },
    ];

    return (
        <div className="overflow-x-auto">
            <Table>
                <TableHeader>
                    <TableRow>
                        {columns.map((col) => (
                            <TableHead key={col.key}>{col.title}</TableHead>
                        ))}
                        <TableHead>Action</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {users.map((user: User) => (
                        <TableRow key={user.id}>
                            <TableCell>{user.full_name}</TableCell>
                            <TableCell>{user.email}</TableCell>
                            <TableCell>{user.gender}</TableCell>
                            <TableCell>
                                {new Date(
                                    user.day_of_birth,
                                ).toLocaleDateString()}
                            </TableCell>
                            <TableCell>
                                {new Date(user.created_at).toLocaleDateString()}
                            </TableCell>
                            <TableCell>
                                {new Date(user.updated_at).toLocaleDateString()}
                            </TableCell>
                            <TableCell>
                                <div className="flex space-x-2">
                                    <Button
                                        variant="default"
                                        //   onClick={showDialog(user.id)}
                                    >
                                        Edit
                                    </Button>
                                    <DeleteItem id={user.id} type="user" />
                                </div>
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </div>
    );
}
