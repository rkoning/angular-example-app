import { Injectable } from '@angular/core';
import { BehaviorSubject, of, Observable } from 'rxjs';
import { Thread } from '@models/thread.model';
import { Comment } from '@models/comment.model';
import { ThreadRoutesService } from '@routes/thread-routes.service';
import { mergeMap, catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ThreadService {

  private threads = new BehaviorSubject<Thread[]>([]);
  Threads = this.threads.asObservable();

  private comments = new BehaviorSubject<Comment[]>([]);
  Comments = this.comments.asObservable();

  constructor(private threadRoutesService: ThreadRoutesService) { }

  public listThreads = (): Observable<Thread[]> => {
    return this.threadRoutesService.list().pipe(
      mergeMap((res) => {
        this.threads.next(res.data);
        return of(res.data);
      }),
      catchError((err) => of(err))
    );
  }

  public addThread = (thread: Thread): Observable<Thread> => {
    return this.threadRoutesService.add(thread).pipe(
      mergeMap((res) => {
        const threads = this.threads.value;
        thread.id = res.data.InsertedID;
        threads.push(thread);
        this.threads.next(threads);
        return of(thread);
      }),
      catchError((err) => of(err))
    );
  }

  public deleteThread = (id: string): Observable<Thread> => {
    return this.threadRoutesService.delete(id).pipe(
      mergeMap((res) => {
        const threads = this.threads.value;
        const idx = threads.findIndex((thread) => thread.id === id);
        const removed = threads[idx];
        threads.splice(idx, 1);
        this.threads.next(threads);
        return of(removed);
      }),
      catchError((err) => of(err))
    );
  }

  public getThreadById = (id: string): Thread => {
    return this.threads.value.find((t) => t.id === id);
  }

  public getThreadComments = (id: string): Observable<Comment[]> => {
    return this.threadRoutesService.getComments(id).pipe(
      mergeMap((res) => {
        this.comments.next(res.data);
        return of(res.data);
      }),
      catchError((err) => of(err))
    );
  }

  public addComment = (threadId: string, comment: Comment): Observable<Comment> => {
    return this.threadRoutesService.addComment(threadId, comment).pipe(
      mergeMap((res) => {
        const comments = this.comments.value;
        comment.id = res.data.InsertedID;
        comments.push(comment);
        this.comments.next(comments);
        return of(comment);
      }),
      catchError((err) => of(err))
    );
  }

  public deleteComment = (threadId: string, commentId: string): Observable<Comment> => {
    return this.threadRoutesService.deleteComment(threadId, commentId).pipe(
      mergeMap((res) => {
        const comments = this.comments.value;
        const idx = comments.findIndex((comment) => comment.id === commentId);
        const removed = comments[idx];
        comments.splice(idx, 1);
        this.comments.next(comments);
        return of(res.data);
      }),
      catchError((err) => of(err))
    );
  }
}
