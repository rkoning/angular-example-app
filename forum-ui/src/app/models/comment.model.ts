export class Comment {
  id: string;
  userId: string;
  parentId: string;
  text: string;

  constructor(id: string, userId: string, parentId: string, text: string) {
    this.id = id;
    this.userId = userId;
    this.parentId = parentId;
    this.text = text;
  }
}
