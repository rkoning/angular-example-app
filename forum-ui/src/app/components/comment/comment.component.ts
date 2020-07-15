import { Component, OnInit, Input } from '@angular/core';
import { Comment } from '@models/comment.model';
import { ThreadService } from '@services/thread.service';

@Component({
  selector: 'app-comment',
  templateUrl: './comment.component.html',
  styleUrls: ['./comment.component.scss']
})
export class CommentComponent implements OnInit {
  @Input() comment: Comment;

  constructor(private threadService: ThreadService) { }

  ngOnInit(): void {
  }

  public deleteComment = (): void => {
    this.threadService.deleteComment(this.comment.parentId, this.comment.id).subscribe((res) => {});
  }
}
