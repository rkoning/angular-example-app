import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BaseRoutesService } from './base-routes.service';
import { Comment } from '@models/comment.model';

@Injectable({
  providedIn: 'root'
})
export class CommentRoutesService extends BaseRoutesService<Comment> {

  public baseRoute = 'comments';

  constructor(protected http: HttpClient) { super(http); }
}
