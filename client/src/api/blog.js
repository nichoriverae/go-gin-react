
export class BlogApi {

  constructor({ fetch, apiString }) {
    this.fetch = fetch;
    this.apiString = apiString;
  }

  getAll = () => this.fetch(this.apiString);

  getById = (id) => {
    let url = new URL(`${this.apiString}/${id}`);
    return this.fetch(url)
  };

  create = (data) => this.fetch(this.apiString, {
    method: "POST",
    body: {
      title: data.title,
      content: data.content,
    }
  });

  delete = (id) => {
    const url = new URL(`${this.apiString}/${id}`);
    return this.fetch(url, {
      method: "DELETE",
    })
  };
}