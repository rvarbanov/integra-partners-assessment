import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { delay } from 'rxjs/operators';

export interface User {
  id: number;
  user_name: string;
  first_name: string;
  last_name: string;
  email: string;
  user_status: string;
  department: string;
}

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private apiUrl = 'http://localhost:8080/api';
  
  private mockUsers: User[] = [
    {
      id: 1,
      user_name: 'johndoe',
      first_name: 'John',
      last_name: 'Doe',
      email: 'john.doe@example.com',
      user_status: 'active',
      department: 'IT'
    },
    {
      id: 2,
      user_name: 'janesmith',
      first_name: 'Jane',
      last_name: 'Smith',
      email: 'jane.smith@example.com',
      user_status: 'active',
      department: 'HR'
    },
    {
      id: 3,
      user_name: 'bobjohnson',
      first_name: 'Bob',
      last_name: 'Johnson',
      email: 'bob.johnson@example.com',
      user_status: 'inactive',
      department: 'Sales'
    },
    {
      id: 4,
      user_name: 'alicebrown',
      first_name: 'Alice',
      last_name: 'Brown',
      email: 'alice.brown@example.com',
      user_status: 'active',
      department: 'Marketing'
    },
    {
      id: 5,
      user_name: 'charliewilson',
      first_name: 'Charlie',
      last_name: 'Wilson',
      email: 'charlie.wilson@example.com',
      user_status: 'terminated',
      department: 'Finance'
    }
  ];

  constructor(private http: HttpClient) {}

  getUsers(): Observable<User[]> {
    // Simulate API call with delay
    return of(this.mockUsers).pipe(delay(500));
  }

  addUser(user: User): Observable<User> {
    return of({ ...user, id: this.mockUsers.length + 1 }).pipe(delay(500));
  }

  updateUser(user: User): Observable<User> {
    return of(user).pipe(delay(500));
  }

  deleteUser(id: number): Observable<void> {
    return of(void 0).pipe(delay(500));
  }
} 