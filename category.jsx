class CategoryManager {
  constructor() {
    this.departments = JSON.parse(localStorage.getItem("departments")) || [
      "Law Department",
      "Business English Department",
      "Japanese Department",
      "French Department",
    ];
    this.programs = JSON.parse(localStorage.getItem("programs")) || [
      "High Quality",
      "Regular",
      "Talented Bachelor",
      "Advanced Program",
    ];
    this.statuses = JSON.parse(localStorage.getItem("statuses")) || [
      "Studying",
      "Graduated",
      "Dropped Out",
      "Temporarily Stopped",
    ];
  }

  addCategory(type, value) {
    if (!value.trim()) return false;

    const list = this[type + "s"];
    if (list.includes(value)) return false;

    list.push(value);
    this.saveToStorage(type);
    return true;
  }

  removeCategory(type, value) {
    const list = this[type + "s"];
    const index = list.indexOf(value);
    if (index > -1) {
      list.splice(index, 1);
      this.saveToStorage(type);
      return true;
    }
    return false;
  }

  saveToStorage(type) {
    localStorage.setItem(type + "s", JSON.stringify(this[type + "s"]));
    this.updateFormSelects();
  }

  updateFormSelects() {
    document.getElementById("department").innerHTML = this.departments
      .map((dept) => `<option value="${dept}">${dept}</option>`)
      .join("");

    document.getElementById("program").innerHTML = this.programs
      .map((prog) => `<option value="${prog}">${prog}</option>`)
      .join("");

    document.getElementById("status").innerHTML = this.statuses
      .map((status) => `<option value="${status}">${status}</option>`)
      .join("");
  }

  displayCategories() {
    document.getElementById("departments-list").innerHTML = this.departments
      .map((dept) => this.createCategoryItem(dept, "department"))
      .join("");

    document.getElementById("programs-list").innerHTML = this.programs
      .map((prog) => this.createCategoryItem(prog, "program"))
      .join("");

    document.getElementById("statuses-list").innerHTML = this.statuses
      .map((status) => this.createCategoryItem(status, "status"))
      .join("");
  }

  createCategoryItem(value, type) {
    return `
            <div class="category-item">
                <span class="category-name" data-value="${value}">${value}</span>
                <div class="btn-group">
                    <button class="update-btn" onclick="updateCategory('${type}', '${value}')">Edit</button>
                    <button class="delete-btn" onclick="deleteCategory('${type}', '${value}')">Delete</button>
                </div>
            </div>
        `;
  }
}

class CategoryManagerWithLogging extends CategoryManager {
  addCategory(type, value) {
    const success = super.addCategory(type, value);
    if (success) {
      logManager.addLog("Add Category", {
        type: type,
        value: value,
      });
    }
    return success;
  }

  removeCategory(type, value) {
    const success = super.removeCategory(type, value);
    if (success) {
      logManager.addLog("Remove Category", {
        type: type,
        value: value,
      });
    }
    return success;
  }
}

const categoryManager = new CategoryManagerWithLogging();

function addCategory(type) {
  const input = document.getElementById(`new-${type}`);
  const value = input.value;
  if (categoryManager.addCategory(type, value)) {
    input.value = "";
    categoryManager.displayCategories();
  } else {
    alert("This category already exists or is invalid!");
  }
}

function deleteCategory(type, value) {
  if (confirm(`Are you sure you want to delete "${value}"?`)) {
    if (categoryManager.removeCategory(type, value)) {
      categoryManager.displayCategories();
    }
  }
}

function updateCategory(type, oldValue) {
  const categoryItem = document.querySelector(
    `.category-name[data-value="${oldValue}"]`
  );
  if (!categoryItem) return;

  const parent = categoryItem.parentElement;
  const btnGroup = parent.querySelector(".btn-group");

  const updateBtn = btnGroup.querySelector(".update-btn");
  const deleteBtn = btnGroup.querySelector(".delete-btn");
  updateBtn.style.display = "none";
  deleteBtn.style.display = "none";

  const input = document.createElement("input");
  input.type = "text";
  input.value = oldValue;
  input.classList.add("edit-input");

  const confirmBtn = document.createElement("button");
  confirmBtn.textContent = "Confirm";
  confirmBtn.classList.add("confirm-btn");

  const cancelBtn = document.createElement("button");
  cancelBtn.textContent = "Cancel";
  cancelBtn.classList.add("cancel-btn");

  function confirmEdit() {
    const newValue = input.value.trim();
    if (!newValue || newValue === oldValue) {
      cancelEdit();
      return;
    }

    if (confirm(`Do you want to change "${oldValue}" to "${newValue}"?`)) {
      categoryManager.removeCategory(type, oldValue);
      categoryManager.addCategory(type, newValue);
      categoryManager.displayCategories();
    } else {
      cancelEdit();
    }
  }

  function cancelEdit() {
    parent.replaceChild(categoryItem, input);
    btnGroup.replaceChild(updateBtn, confirmBtn);
    btnGroup.replaceChild(deleteBtn, cancelBtn);
    updateBtn.style.display = "inline-block";
    deleteBtn.style.display = "inline-block";
  }

  input.addEventListener("keydown", (e) => {
    if (e.key === "Enter") confirmEdit();
    if (e.key === "Escape") cancelEdit();
  });

  confirmBtn.addEventListener("click", confirmEdit);
  cancelBtn.addEventListener("click", cancelEdit);

  parent.replaceChild(input, categoryItem);

  btnGroup.appendChild(confirmBtn);
  btnGroup.appendChild(cancelBtn);

  input.focus();
}

document.querySelectorAll(".tab-button").forEach((button) => {
  button.addEventListener("click", () => {
    document
      .querySelectorAll(".tab-button")
      .forEach((b) => b.classList.remove("active"));
    document
      .querySelectorAll(".tab-content")
      .forEach((c) => c.classList.remove("active"));
    button.classList.add("active");
    document.getElementById(button.dataset.tab).classList.add("active");

    if (button.dataset.tab === "list") {
      displayStudents();
    } else if (button.dataset.tab === "categories") {
      categoryManager.displayCategories();
    }
  });
});

categoryManager.updateFormSelects();
categoryManager.displayCategories();
