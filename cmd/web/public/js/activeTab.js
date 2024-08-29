const availableTab = document.getElementById("available-tab");
const collectedTab = document.getElementById("collected-tab");
const chatTab = document.getElementById("chat-tab");

const tabs = [availableTab, collectedTab, chatTab];

tabs.forEach((tab) => {
  tab.addEventListener("click", function () {
    const selectedTab = document.querySelector(".active");
    selectedTab.classList.remove("active");
    this.classList.add("active");
  });
});
