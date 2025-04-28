document.addEventListener('DOMContentLoaded', () => {
  const path = window.location.pathname;

  /*
  if (path.includes('projects.html')) {
    loadProjects();
  }

  

  if (path.includes('contact.html')) {
    setupContactForm();
  }

  */
  loadProjects();
  setupContactForm();

});

// Load and display projects
function loadProjects() {
  fetch('https://portfolio-04d9.onrender.com/api/projects')
    .then(res => res.json())
    .then(projects => {
      const projectList = document.getElementById('project-list');
      projects.forEach(project => {
        const card = document.createElement('div');
        card.className = 'project-card';
        card.innerHTML = `
          <img src="${project.image_url}" alt="${project.name}" />
          <h3>${project.name}</h3>
          <p>${project.description}</p>
          <p><strong>${project.date_range}</strong></p>
          <a href="${project.github_url}" target="_blank">GitHub</a>
        `;
        projectList.appendChild(card);
      });
    })
    .catch(err => {
      console.error('Failed to load projects:', err);
    });
}




// Handle contact form (no backend yet)
function setupContactForm() {
  const form = document.querySelector('#contact-form');

  form.addEventListener('submit', async e => {
    e.preventDefault();

    const name = form.querySelector('#name').value.trim();
    const email = form.querySelector('#email').value.trim();
    const message = form.querySelector('#message').value.trim();

    if (!name || !email || !message) {
      alert('Please fill in all fields.');
      return;
    }

    console.log("Preparing to send request...")

    try {
      const res = await fetch("https://portfolio-04d9.onrender.com/api/contact", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, email, message })
      });

      const result = await res.json();

      console.log(result)

      if (res.ok) {
        alert("Message sent! I look forward to speaking with you!");
        form.reset();
      } else {
        alert("Failed to send message.");
      }
    } catch (error) {
      console.error("Error sending message:", error);
      alert("Something went wrong. Please try again later.");
    }
  });
}

