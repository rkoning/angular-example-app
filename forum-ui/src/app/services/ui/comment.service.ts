import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { Comment } from '@models/comment.model';

@Injectable({
  providedIn: 'root'
})
export class CommentService {

  private comments = new BehaviorSubject<Comment[]>([]);
  Comments = this.comments.asObservable();

  constructor() { }

  public setComments = (comments: Comment[]): void => {
    this.comments.next(comments);
  }

  public addComment = (comment: Comment): void => {
    const comments = this.comments.value;
    comments.push(comment);
    this.comments.next(comments);
  }
}
