export const getUsers = async () => {
    await new Promise((resolve) => setTimeout(resolve, 3000));
    const res = await fetch("http://localhost:8080/api/v1/admin/users", {
        cache: "no-cache",
    });
    if (!res.ok) {
        throw new Error("Failed to fetch users.");
    }
    return res.json();
};

export const deleteUser = async (id: string) => {
    await new Promise((resolve) => setTimeout(resolve, 3000));
    const res = await fetch(`http://localhost:8080/api/v1/admin/users/${id}`, {
        method: "DELETE",
        cache: "no-cache",
    });
    if (!res.ok) {
        throw new Error("Failed to delete user.");
    }
    return res.json();
};

export const getProducts = async ({
    page = 1,
    limit = 5,
    search = "",
    sort = "desc",
}: {
    page: number;
    limit: number;
    search: string;
    sort: string;
}) => {
    await new Promise((resolve) => setTimeout(resolve, 1000));
    const query = new URLSearchParams({
        page: page.toString(),
        limit: limit.toString(),
        search,
        sort,
    });

    const res = await fetch(
        `http://localhost:8080/api/v1/admin/products?${query.toString()}`,
    );

    if (!res.ok) {
        return "Failed to fetch products.";
    }

    return res.json();
};

export const deleteProduct = async (id: string) => {
    await new Promise((resolve) => setTimeout(resolve, 3000));
    const res = await fetch(
        `http://localhost:8080/api/v1/admin/products/${id}`,
        {
            method: "DELETE",
            cache: "no-cache",
        },
    );
    if (!res.ok) {
        throw new Error("Failed to delete product.");
    }
    return res.json();
};
