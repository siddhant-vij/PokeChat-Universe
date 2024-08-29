document.addEventListener("DOMContentLoaded", function () {
  const tabs = document.querySelectorAll(
    "#available-tab, #collected-tab, #chat-tab"
  );

  tabs.forEach((tab) => {
    if (tab.dataset.disabled === "true") {
      tab.addEventListener("click", function (event) {
        event.preventDefault();
        event.stopImmediatePropagation();
        return false;
      });
    } else {
      tab.addEventListener("click", function () {
        const selectedTab = document.querySelector(".active");
        selectedTab.classList.remove("active");
        this.classList.add("active");
      });
    }
  });
});
