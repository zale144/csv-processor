CREATE TABLE IF NOT EXISTS public.user_batch(
     user_id INTEGER NOT NULL UNIQUE REFERENCES public.user(id),
     saved_to_crm BOOL DEFAULT false
);
