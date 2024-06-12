function addCategory() {
    const categoryName = document.getElementById("cat").value.trim();
    const subjectName  = document.getElementById("sub").value.trim();

    if (categoryName === "") {
        alert("il faut remplir les champs");
        return;
    }

    const donnees = new URLSearchParams();
    donnees.append("cat", categoryName);
    donnees.append("sub", subjectName);

    fetch("/category", {
        method: "POST",
        body: donnees
    }).then(response => {
        if (!response.ok) {
            throw new Error("Erreur HTTP : " + response.status);
        }
        return response.json();
    }).then(donnees => {
        console.log("Categoire ajoutée avec succès", donnees);
        location.reload();
    }).catch(error => {
        console.error("Erreur :", error);
        alert("Erreur lors de l'ajout de la catégorie");
    });

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

    const newInputSubject = document.createElement("input");
    newInputSubject.setAttribute("type", "text");
    newInputSubject.classList.add("sub-btn");

    const newButtonAdd = document.createElement("button");
    newButtonAdd.type = "button";
    newButtonAdd.textContent = "Ajouter";
    newButtonAdd.classList.add("add-btn");

    newButtonAdd.addEventListener("click", function() {
        const subjectName = newInputSubject.value.trim();
        if (subjectName === "") {
            alert("il faut remplir les champs");
            return;
        }

        const newSubject = document.createElement("li");
        const newSubjectLink = document.createElement("a");
        newSubjectLink.setAttribute("href", "#");
        newSubjectLink.textContent = subjectName;
        newSubject.appendChild(newSubjectLink);
        subjectList.appendChild(newSubject);
        newInputSubject.value = "";
    });


    categoryContainer.appendChild(subjectList);
    categoryContainer.appendChild(newInputSubject);
    categoryContainer.appendChild(newButtonAdd);
    newCategory.appendChild(categoryTitle);
    newCategory.appendChild(categoryContainer);

    const categoryList = document.querySelector(".container");
    const secondCategory = categoryList.children[1];
    categoryList.insertBefore(newCategory, secondCategory);

    document.getElementById("cat").value = "";
    document.getElementById("sub").value = "";
}

document.getElementById("addCategory").addEventListener("click", addCategory);
