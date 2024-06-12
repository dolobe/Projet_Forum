var categoryLink = document.getElementById("categoryLink");
var categoryDropdown = document.getElementById("categoryDropdown");

function showCategories() {
    categoryDropdown.style.display = "block";
}

categoryDropdown.addEventListener("mouseleave", function(event) {
    if (!categoryDropdown.contains(event.relatedTarget)) {
        hideCategories();
    }
});

function hideCategories() {
    setTimeout(function() {
        if (!categoryDropdown.contains(document.querySelector(":hover"))) {
            categoryDropdown.style.display = "none";
        }
    }, 300);
}

categoryLink.addEventListener("mouseover", showCategories);
categoryDropdown.addEventListener("mouseleave", hideCategories);


