function addCategory() {
    const categoryName = document.getElementById("cat").value.trim();
    const subjectName  = document.getElementById("sub").value.trim();

    const newCategory = document.createElement("div");
    newCategory.classList.add("category-item");

    const categoryTitle = document.createElement("div");
    categoryTitle.classList.add("category-title");
    categoryTitle.innerHTML = '<a href="#">' + categoryName + '</a>';

    const categoryContainer = document.createElement("div");
    categoryContainer.classList.add("category-container");

    const subjectList = document.createElement("ul");
    subjectList.setAttribute("id", "Subject-container");

    if (subjectList) {
        const subject1 = document.createElement("li");
        const subject2 = document.createElement("a");
        subject2.setAttribute("href", "#");
        subject2.textContent = subjectName;
        subject1.appendChild(subject2);
        subjectList.appendChild(subject1);
    }


    categoryContainer.appendChild(subjectList);
    newCategory.appendChild(categoryTitle);
    newCategory.appendChild(categoryContainer);

    const categoryList = document.querySelector(".container");
    categoryList.appendChild(newCategory);

    document.getElementById("cat").value = "";
    document.getElementById("sub").value = "";
}

document.getElementById("addCategory").addEventListener("click", addCategory);
