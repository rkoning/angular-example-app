import { Component, OnInit, Input } from '@angular/core';
import { Thread } from '@models/thread.model';
import { ThreadRoutesService } from '@routes/thread-routes.service';
import { Router, ActivatedRoute } from '@angular/router';
import { ThreadService } from '@services/thread.service';
import { CommentRoutesService } from '@routes/comment-routes.service';
import { CommentService } from '@services/comment.service';
import { Comment } from '@models/comment.model';

@Component({
  selector: 'app-thread',
  templateUrl: './thread.component.html',
  styleUrls: ['./thread.component.scss']
})
export class ThreadComponent implements OnInit {
  @Input() thread: Thread;

  public comments: Comment[];
  public showSpinner = true;
  public showCommentForm = false;

  constructor(private route: ActivatedRoute, private commentService: CommentService,
              private commentRoutes: CommentRoutesService, private threadService: ThreadService,
              private threadRoutesService: ThreadRoutesService) { }

  ngOnInit(): void {
    this.commentService.Comments.subscribe((comments) => this.comments = comments);
    this.threadRoutesService.getComments(this.thread.id).subscribe((res) => {
      this.commentService.setComments(res.data);
      this.showSpinner = false;
    });
  }
}
