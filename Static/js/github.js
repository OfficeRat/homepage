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

    const column = document.createElement('div');
    column.classList.add('column', 'is-one-third');

    const box = document.createElement('div');
    box.classList.add('is-clickable','box','has-text-white');
    box.style.backgroundColor = "#354D66";

    const repoName = document.createElement('h3');
    repoName.classList.add('is-capitalized','title', 'is-size-4','has-text-white'); 
    repoName.textContent = repo.name;

    const repoDescription = document.createElement('p');
    repoDescription.textContent = repo.description || 'No description available.';

    box.appendChild(repoName);
    box.appendChild(repoDescription);

    box.addEventListener('click', () => {
        window.open(repo.html_url, '_blank');
    });

    column.appendChild(box);

    repoList.appendChild(column);
}


