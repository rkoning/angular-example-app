export class Thread {
  id: string;
  userId: string;
  title: string;
  text: string;

  constructor(id, userId, title, text) {
    this.id = id;
    this.userId = userId;
    this.title = title;
    this.text = text;
  }
}
