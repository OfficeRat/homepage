fetch(`https://api.github.com/users/OfficeRat/repos`)
  .then(response => response.json())
  .then(repositories => {
    repositories.forEach(repository => {
      createRepoElement(repository);
    });
  })
  .catch(error => console.error('Error fetching repositories:', error));

function createRepoElement(repo) {
    const repoList = document.getElementById('repo-list');
  
    const repoElement = document.createElement('div');
    repoElement.classList.add('repo');
  
    const repoName = document.createElement('h3');
    repoName.textContent = repo.name;
  
    const repoDescription = document.createElement('p');
    repoDescription.textContent = repo.description || 'No description available.';

    repoElement.addEventListener('click', () => {
      window.open(repo.html_url, '_blank');
    });
  
    repoElement.appendChild(repoName);
    repoElement.appendChild(repoDescription);
  
    repoList.appendChild(repoElement);
  }