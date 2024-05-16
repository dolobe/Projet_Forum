function addCategory() {
    const categoryName = document.getElementById("cat").value.trim();

    const newCategory = document.createElement("a");
    newCategory.classList.add("category-item");
    newCategory.setAttribute("href", "#");

    const categoryTitle = document.createElement("div");
    categoryTitle.classList.add("category-title");
    categoryTitle.textContent = categoryName;

    const categoryContainer = document.createElement("div");
    categoryContainer.classList.add("category-container");

    const subjectList = document.createElement("ul");
    subjectList.setAttribute("id", "Subject-container");

    categoryContainer.appendChild(subjectList);
    newCategory.appendChild(categoryTitle);
    newCategory.appendChild(categoryContainer);

    const categoryList = document.querySelector(".container");
    categoryList.appendChild(newCategory);

    document.getElementById("cat").value = "";
}

document.getElementById("addCategory").addEventListener("click", addCategory);
