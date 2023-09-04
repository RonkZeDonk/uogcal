ALTER TABLE IF EXISTS public.section_meeting DROP CONSTRAINT IF EXISTS section_meeting_uq;
ALTER TABLE IF EXISTS public.section_meeting DROP CONSTRAINT IF EXISTS course_section_section_meeting_fk;

ALTER TABLE IF EXISTS public.section_meeting
    ADD CONSTRAINT course_section_section_meeting_fk FOREIGN KEY (code)
    REFERENCES public.course_section (code) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;