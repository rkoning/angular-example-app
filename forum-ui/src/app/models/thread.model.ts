export class Thread {
  id: string;
  user: string;
  title: string;
  text: string;

  constructor(id, user, title, text) {
    this.id = id;
    this.user = user;
    this.title = title;
    this.text = text;
  }
}
