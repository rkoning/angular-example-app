import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { Thread } from '@models/thread.model';

@Injectable({
  providedIn: 'root'
})
export class ThreadService {

  private threads = new BehaviorSubject<Thread[]>([]);
  Threads = this.threads.asObservable();

  constructor() { }

  public addThread = (thread: Thread): void => {
    const threads = this.threads.value;
    thread.id = threads.length.toString();
    threads.push(thread);
    this.threads.next(threads);
  }

  public getThreadById = (id: string): Thread => {
    console.log(this.threads.value);
    console.log(id);
    return this.threads.value.find((t) => t.id === id);
  }
}
